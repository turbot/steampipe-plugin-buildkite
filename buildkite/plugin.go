package buildkite

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-buildkite",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"buildkite_agent": 		  tableBuildkiteAgent(ctx),
			"buildkite_build":        tableBuildkiteBuild(ctx),
			"buildkite_organization": tableBuildkiteOrganization(ctx),
			"buildkite_pipeline":     tableBuildkitePipeline(ctx),
			"buildkite_team":         tableBuildkiteTeam(ctx),
			"buildkite_user":         tableBuildkiteUser(ctx),
		},
	}
	return p
}
