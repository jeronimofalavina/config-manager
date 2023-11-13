package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	model "github.com/jeronimofalavina/config-manager/api/cmd/models"

	"github.com/gin-gonic/gin"
)

var configs = []model.Config{}

func ListConfigs(c *gin.Context) {
	c.JSON(http.StatusOK, configs)
}

func CreateConfigs(c *gin.Context) {
	var newConfig []model.Config
	if err := c.BindJSON(&newConfig); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for index, existConf := range newConfig {
		_, i, _ := getConfigByName(existConf.Name)
		if i == -1 {
			configs = append(configs, newConfig[index])
			c.JSON(http.StatusCreated, newConfig[index])
		} else if i != -1 {
			c.JSON(http.StatusContinue, gin.H{"message": "config " + newConfig[index].Name + " already exists."})
		}
	}

}

func GetConfig(c *gin.Context) {
	name := c.Param("name")
	config, _, err := getConfigByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "config not found."})
		return
	}

	c.JSON(http.StatusOK, config)
}

func UpdateConfig(c *gin.Context) {
	name := c.Param("name")
	updatedConfig, index, err := getConfigByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "config not found."})
		return
	}

	var updatedData model.Config

	if err := c.BindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// TODO: check with the new value already exists in the list of configs, to prevent duplicated values

	if c.Request.Method == http.MethodPatch {
		if updatedData.Name != "" {
			updatedConfig.Name = updatedData.Name
		}
		if updatedData.Metadata != nil {
			updatedConfig.Metadata = updatedData.Metadata
		}
	} else if c.Request.Method == http.MethodPut {
		updatedConfig.Name = updatedData.Name
		updatedConfig.Metadata = updatedData.Metadata
	} else {
		c.JSON(http.StatusMethodNotAllowed, gin.H{"message": "Unsupported HTTP method."})
		return
	}

	configs[index] = *updatedConfig

	c.JSON(http.StatusOK, updatedConfig)
}

func QueryConfigs(c *gin.Context) {
	queryParam := c.Request.URL.Query()
	var matchConfigs []model.Config
	var keys []string
	var valueArrays [][]string

	for key, values := range queryParam {
		keys = append(keys, key)
		valueArrays = append(valueArrays, values)
	}

	lastKey := keys[len(keys)-1]
	aux := strings.Split(lastKey, ".")
	lastKey = aux[len(aux)-1]

	aux = valueArrays[len(valueArrays)-1]
	value := aux[len(valueArrays)-1]

	for _, conf := range configs {
		if nestedMapValueMatches(conf.Metadata, lastKey, value) {
			matchConfigs = append(matchConfigs, conf)
		}
	}
	if len(matchConfigs) == 0 {
		c.JSON(http.StatusNotFound, nil)
	}
	c.JSON(http.StatusOK, matchConfigs)

}

func nestedMapValueMatches(metadata map[string]interface{}, key string, value string) bool {
	for k, v := range metadata {
		if k == key {
			if fmt.Sprintf("%v", v) == value {
				return true
			}
		}

		if nestedMap, ok := v.(map[string]interface{}); ok {
			if nestedMapValueMatches(nestedMap, key, value) {
				return true
			}
		}
	}

	return false
}

func DeleteConfig(c *gin.Context) {
	name := c.Param("name")
	_, index, err := getConfigByName(name)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Config not found."})
		return
	}

	configs = append(configs[:index], configs[index+1:]...)

	c.JSON(http.StatusOK, gin.H{"message": "Config deleted successfully."})
}

func getConfigByName(name string) (*model.Config, int, error) {
	for i, conf := range configs {
		if conf.Name == name {
			return &configs[i], i, nil
		}
	}

	return nil, -1, errors.New("config not found")
}
