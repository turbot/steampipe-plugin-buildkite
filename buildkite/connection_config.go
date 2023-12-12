package buildkite

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type buildkiteConfig struct {
	Token *string `hcl:"token"`
}

func ConfigInstance() interface{} {
	return &buildkiteConfig{}
}

// GetConfig :: retrieve and cast connection config from query data
func GetConfig(connection *plugin.Connection) buildkiteConfig {
	if connection == nil || connection.Config == nil {
		return buildkiteConfig{}
	}
	config, _ := connection.Config.(buildkiteConfig)
	return config
}
