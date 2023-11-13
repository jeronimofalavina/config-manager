package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"HFtest-platform-jeronimofalavina-tst/cmd/handler"
	"HFtest-platform-jeronimofalavina-tst/cmd/models"

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
	models.Configs = []models.Config{
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
	r.GET("/configs", handler.ListConfigs)
	req, _ := http.NewRequest("GET", "/configs", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var responseConfigs []models.Config
	err := json.Unmarshal(w.Body.Bytes(), &responseConfigs)
	if err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	assert.Equal(t, models.Configs, responseConfigs)
	assert.Equal(t, http.StatusOK, w.Code)
}
