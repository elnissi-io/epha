package config

import (
	"fmt"
)

var Configs = map[string]ResourceConfig{
    // key-value pairs of configurations
}

type ResourceConfig struct {
    Template     string
    RequiredVars []string
}


type ConfigFactory struct {
	resourceTemplates map[string]string
}

func NewConfigFactory() *ConfigFactory {
	return &ConfigFactory{
		resourceTemplates: make(map[string]string),
	}
}

func (f *ConfigFactory) RegisterResource(name, template string) {
	f.resourceTemplates[name] = template
}

func (f *ConfigFactory) GetResourceConfig(name string) (string, error) {
	template, exists := f.resourceTemplates[name]
	if !exists {
		return "", fmt.Errorf("no config available for resource: %s", name)
	}
	return template, nil
}
