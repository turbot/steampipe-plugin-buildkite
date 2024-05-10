package buildkite

import (
	"context"

	bkapi "github.com/buildkite/go-buildkite/v3/buildkite"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBuildkiteOrganization(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "buildkite_organization",
		Description: "Organizations in the Buildkite account.",
		List: &plugin.ListConfig{
			Hydrate: listOrganization,
		},
		Get: &plugin.GetConfig{
			Hydrate: getOrganization,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "slug"},
			},
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "Slug of the organization."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the organization."},
			// Other columns
			{Name: "agents_url", Type: proto.ColumnType_STRING, Description: "Agents URL for the organization."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timeToRfc3339), Description: "Time when the organization was created."},
			{Name: "emojis_url", Type: proto.ColumnType_STRING, Description: "Emojis URL for the organization."},
			{Name: "graphql_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("GraphQLID"), Description: "GraphQL ID for the organization."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the organization."},
			{Name: "pipelines_url", Type: proto.ColumnType_STRING, Description: "Pipelines URL for the organization."},
			{Name: "repository", Type: proto.ColumnType_STRING, Description: "Repository for the organization."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL for the organization."},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "Web URL for the organization."},
		}),
	}
}

func listOrganization(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_organization.listOrganization", "connection_error", err)
		return nil, err
	}

	opts := &bkapi.OrganizationListOptions{
		ListOptions: bkapi.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for {
		organizations, resp, err := conn.Organizations.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("buildkite_organization.listOrganization", "query_error", err, "resp", resp)
			return nil, err
		}
		for _, i := range organizations {
			d.StreamListItem(ctx, i)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, nil
}

func getOrganization(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_organization.getOrganization", "connection_error", err)
		return nil, err
	}

	slug := d.EqualsQuals["slug"].GetStringValue()

	organization, resp, err := conn.Organizations.Get(slug)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_organization.getOrganization", "query_error", err, "resp", resp, "slug", slug)
		return nil, err
	}

	return organization, nil
}
