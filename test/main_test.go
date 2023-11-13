package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	setupTestData()
	return router
}

func setupTestData() {
	configs = []Config{
		{
			Name: "datacenter-1",
			Metadata: map[string]interface{}{
				"monitoring": map[string]interface{}{
					"enabled": "true",
				},
				"limits": map[string]interface{}{
					"cpu": map[string]interface{}{
						"enabled": "false",
						"value":   "300m",
					},
				},
			},
		},
		{
			Name: "datacenter-2",
			Metadata: map[string]interface{}{
				"monitoring": map[string]interface{}{
					"enabled": "false",
				},
				"limits": map[string]interface{}{
					"cpu": map[string]interface{}{
						"enabled": "true",
						"value":   "250m",
					},
				},
			},
		},
		{
			Name: "burger-nutrition",
			Metadata: map[string]interface{}{
				"allergens": map[string]interface{}{
					"eggs":    "true",
					"nuts":    "false",
					"seafood": "false",
				},
				"calories": 230.0,
				"carbohydrates": map[string]interface{}{
					"dietary-fiber": "4g",
					"sugars":        "1g",
				},
				"fats": map[string]interface{}{
					"saturated-fat": "0g",
					"trans-fat":     "1g",
				},
			},
		},
	}
}

func TestListConfigs(t *testing.T) {

	r := SetUpRouter()
	r.GET("/configs", listConfigs)
	req, _ := http.NewRequest("GET", "/configs", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseConfigs []Config
	err := json.Unmarshal(w.Body.Bytes(), &responseConfigs)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, configs, responseConfigs)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateConfigs(t *testing.T) {
	r := SetUpRouter()
	r.POST("/configs", createConfigs)
	conf := []Config{
		{
			Name: "datacenter-4",
			Metadata: map[string]interface{}{
				"monitoring": map[string]interface{}{
					"enabled": "false",
				},
				"limits": map[string]interface{}{
					"cpu": map[string]interface{}{
						"enabled": "false",
						"value":   "300m",
					},
				},
			},
		},
		{
			Name: "datacenter-5",
			Metadata: map[string]interface{}{
				"monitoring": map[string]interface{}{
					"enabled": "false",
				},
				"limits": map[string]interface{}{
					"cpu": map[string]interface{}{
						"enabled": "true",
						"value":   "266m",
					},
				},
			},
		}}

	jsonValue, _ := json.Marshal(conf)
	req, _ := http.NewRequest("POST", "/configs", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestCreateConfigsExists(t *testing.T) {
	r := SetUpRouter()

	r.POST("/configs", createConfigs)

	jsonValue, _ := json.Marshal(configs)
	req, _ := http.NewRequest("POST", "/configs", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusContinue, w.Code)
}

func TestGetConfig(t *testing.T) {
	r := SetUpRouter()

	r.GET("/configs/:name", getConfig)
	testName := "datacenter-1"
	req, _ := http.NewRequest("GET", "/configs/"+testName, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var config Config
	json.Unmarshal(w.Body.Bytes(), &config)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, testName, config.Name)
}

func TestUpdateConfigPUT(t *testing.T) {
	r := SetUpRouter()
	configName := "datacenter-1"
	newConfigName := "datacenter-6"
	r.PUT("/configs/:name", updateConfig)
	updateConf := Config{

		Name: newConfigName,
		Metadata: map[string]interface{}{
			"monitoring": map[string]interface{}{
				"enabled": "false",
			},
			"limits": map[string]interface{}{
				"cpu": map[string]interface{}{
					"enabled": "false",
					"value":   "266m",
				},
			},
		},
	}

	jsonValue, _ := json.Marshal(updateConf)
	req, _ := http.NewRequest("PUT", "/configs/"+configName, bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var config []Config
	json.Unmarshal(w.Body.Bytes(), &config)

	conf, _, _ := getConfigByName(newConfigName)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, updateConf.Name, conf.Name)
}

func TestDeleteConfig(t *testing.T) {
	r := SetUpRouter()
	configName := "datacenter-1"
	r.DELETE("/configs/:name", deleteConfig)

	req, _ := http.NewRequest("DELETE", "/configs/"+configName, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var config []Config
	json.Unmarshal(w.Body.Bytes(), &config)

	_, i, _ := getConfigByName(configName)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, i, -1)
}

func TestQueryConfigs(t *testing.T) {
	r := SetUpRouter()
	r.GET("/search", queryConfigs)
	query := "/search?metadata.monitoring.enabled=true"
	req, _ := http.NewRequest("GET", query, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseConfig []Config
	err := json.Unmarshal(w.Body.Bytes(), &responseConfig)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, http.StatusOK, w.Code)
}
