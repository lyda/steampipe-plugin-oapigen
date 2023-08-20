// Package oapigen generates tables from OpenAPI documents.
package oapigen

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

// Config is a go type to represent the plugin config.
type Config struct {
	Version   int      `cty:"version"`
	Documents []string `cty:"documents" steampipe:"watch"`
	Prefix    *string  `cty:"prefix"`
}

// ConfigSchema defines the config params for the oapigen plugin.
var ConfigSchema = map[string]*schema.Attribute{
	"version": {
		Type: schema.TypeInt,
	},
	"documents": {
		Type: schema.TypeList,
		Elem: &schema.Attribute{Type: schema.TypeString},
	},
	"prefix": {
		Type: schema.TypeString,
	},
}

// ConfigInstance returns a Config pointer.
func ConfigInstance() interface{} {
	return &Config{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) Config {
	if connection == nil || connection.Config == nil {
		return Config{}
	}
	config, _ := connection.Config.(Config)
	if config.Version == 0 {
		config.Version = 3
	}
	return config
}
