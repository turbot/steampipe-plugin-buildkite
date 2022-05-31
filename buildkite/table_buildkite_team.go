package buildkite

import (
	"context"

	bkapi "github.com/buildkite/go-buildkite/v3/buildkite"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableBuildkiteTeam(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "buildkite_team",
		Description: "Teams in the Buildkite account.",
		List: &plugin.ListConfig{
			ParentHydrate: listOrganization,
			Hydrate:       listTeam,
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_slug", Type: proto.ColumnType_STRING, Description: "Organization slug for the team."},
			{Name: "slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.Slug"), Description: "Slug of the team."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.Name"), Description: "Name of the team."},
			// Other columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.ID"), Description: "ID of the team."},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.Description"), Description: "Description of the team."},
			{Name: "privacy", Type: proto.ColumnType_STRING, Transform: transform.FromField("Team.Privacy"), Description: "Privacy setting for the team, e.g. visible, secret."},
			{Name: "default", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Team.Default"), Description: "True if this is the default team."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Team.CreatedAt").Transform(timeToRfc3339), Description: "Time when the team was created."},
			{Name: "created_by", Type: proto.ColumnType_JSON, Transform: transform.FromField("Team.CreatedBy"), Description: "User who created the team."},
		},
	}
}

type teamRow struct {
	bkapi.Team
	OrganizationSlug string `json:"organization_slug"`
}

func listTeam(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_team.listTeam", "connection_error", err)
		return nil, err
	}

	// Convenience
	keyQuals := d.KeyColumnQuals

	var org string
	if h.Item != nil {
		org = *h.Item.(bkapi.Organization).Slug
	} else if keyQuals["organization_slug"] != nil {
		org = keyQuals["organization_slug"].GetStringValue()
	}

	opts := &bkapi.TeamsListOptions{
		ListOptions: bkapi.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for {
		teams, resp, err := conn.Teams.List(org, opts)
		if err != nil {
			plugin.Logger(ctx).Error("buildkite_team.listTeam", "query_error", err, "resp", resp)
			return nil, err
		}
		for _, i := range teams {
			d.StreamListItem(ctx, teamRow{i, org})
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, nil
}
