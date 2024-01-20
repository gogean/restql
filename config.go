// config.go

package restql

import (
	"os"

	"gopkg.in/yaml.v2"
)

// APIConfig represents the structure of the YAML configuration file
type APIConfig struct {
	APIs []API `yaml:"apis"`
}

// API defines the structure for an API object.
type API struct {
	Name        string         `yaml:"name" json:"name"`
	Path        string         `yaml:"path" json:"path"`
	Method      string         `yaml:"method" json:"method"`
	Description string         `yaml:"description" json:"description"`
	Parameters  []APIParameter `yaml:"parameters" json:"parameters"`
}

type APIParameter struct {
	Name        string `yaml:"name" json:"name"`
	Type        string `yaml:"type" json:"type"`
	Description string `yaml:"description" json:"description"`
	Required    bool   `yaml:"required" json:"required"`
}

// readConfig reads the YAML configuration file
func GetAPIConfig(filename string) (*APIConfig, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	config := &APIConfig{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func GetAPIs(config *APIConfig) []API {
	var apis []API

	for _, apiEntry := range config.APIs {
		api := API(apiEntry)
		api.Parameters = make([]APIParameter, len(apiEntry.Parameters))

		for i, paramEntry := range apiEntry.Parameters {
			api.Parameters[i] = APIParameter(paramEntry)
		}

		apis = append(apis, api)
	}

	return apis
}
