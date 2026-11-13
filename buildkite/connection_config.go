package buildkite

import (
	"github.com/turbot/steampipe-plugin-sdk/v6/plugin"
)

type buildkiteConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &buildkiteConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) buildkiteConfig {
	if connection == nil || connection.GetConfig() == nil {
		return buildkiteConfig{}
	}
	config, _ := connection.GetConfig().(buildkiteConfig)
	return config
}
