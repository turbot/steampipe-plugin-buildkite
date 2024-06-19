package buildkite

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-buildkite",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
		},
		DefaultTransform: transform.FromJSONTag().NullIfZero(),
		// User ID against which the access token has generated. (This is unique per connection)
		// User can have access to multiple Organizations so org ID can not be include as connection key quals.
		ConnectionKeyColumns: []plugin.ConnectionKeyColumn{
			{
				Name:    "member_id",
				Hydrate: getMemberId,
			},
		},
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		TableMap: map[string]*plugin.Table{
			"buildkite_agent":        tableBuildkiteAgent(ctx),
			"buildkite_build":        tableBuildkiteBuild(ctx),
			"buildkite_organization": tableBuildkiteOrganization(ctx),
			"buildkite_pipeline":     tableBuildkitePipeline(ctx),
			"buildkite_team":         tableBuildkiteTeam(ctx),
			"buildkite_user":         tableBuildkiteUser(ctx),
		},
	}
	return p
}
