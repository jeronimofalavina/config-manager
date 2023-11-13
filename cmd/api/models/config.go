package models

type Config struct {
	Name     string                 `json:"name"`
	Metadata map[string]interface{} `json:"metadata"`
}

var Configs []Config
