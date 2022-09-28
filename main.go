package main

import (
	"github.com/turbot/steampipe-plugin-buildkite/buildkite"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: buildkite.Plugin})
}
