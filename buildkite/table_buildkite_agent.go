package buildkite

import (
	"context"

	bkapi "github.com/buildkite/go-buildkite/v3/buildkite"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBuildkiteAgent(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "buildkite_agent",
		Description: "Agents in the Buildkite account.",
		List: &plugin.ListConfig{
			ParentHydrate: listOrganization,
			Hydrate:       listAgent,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "organization_slug", Require: plugin.Optional},
				{Name: "name", Require: plugin.Optional},
				{Name: "hostname", Require: plugin.Optional},
				{Name: "version", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate: getAgent,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "organization_slug"},
				{Name: "id"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL").Transform(orgFromURL), Description: "Organization of the agent."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the agent."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the agent."},
			{Name: "ip_address", Type: proto.ColumnType_IPADDR, Description: "IP address for the agent."},
			// Other columns
			// Not returned in get / list calls {Name: "access_token", Type: proto.ColumnType_STRING, Description: "Token for the agent."},
			{Name: "connection_state", Type: proto.ColumnType_STRING, Transform: transform.FromField("ConnectedState"), Description: "State of the connection to the agent."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timeToRfc3339), Description: "Time when the agent was created."},
			{Name: "creator", Type: proto.ColumnType_JSON, Description: "User that created the agent."},
			{Name: "hostname", Type: proto.ColumnType_STRING, Description: "Hostname for the agent."},
			{Name: "job", Type: proto.ColumnType_JSON, Description: "Job running on the agent."},
			{Name: "last_job_finished_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("LastJobFinishedAt").Transform(timeToRfc3339), Description: "Time when the agent finished it's last job."},
			{Name: "meta_data", Type: proto.ColumnType_JSON, Description: "Metadata for the agent."},
			{Name: "priority", Type: proto.ColumnType_INT, Description: "Priority of the agent."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL for the agent."},
			{Name: "user_agent", Type: proto.ColumnType_STRING, Description: "User agent of the agent."},
			{Name: "version", Type: proto.ColumnType_STRING, Description: "Version of the agent."},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "Web URL for the agent."},
		},
	}
}

func listAgent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_agent.listAgent", "connection_error", err)
		return nil, err
	}

	// Convenience
	keyQuals := d.EqualsQuals

	var org string
	if h.Item != nil {
		org = *h.Item.(bkapi.Organization).Slug
	} else if keyQuals["organization_slug"] != nil {
		org = keyQuals["organization_slug"].GetStringValue()
	}

	opts := &bkapi.AgentListOptions{
		ListOptions: bkapi.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}

	// Optional quals
	if keyQuals["name"] != nil {
		opts.Name = keyQuals["name"].GetStringValue()
	}
	if keyQuals["hostname"] != nil {
		opts.Hostname = keyQuals["hostname"].GetStringValue()
	}
	if keyQuals["version"] != nil {
		opts.Version = keyQuals["version"].GetStringValue()
	}

	for {
		agents, resp, err := conn.Agents.List(org, opts)
		if err != nil {
			plugin.Logger(ctx).Error("buildkite_agent.listAgent", "query_error", err, "resp", resp, "opts", opts)
			return nil, err
		}
		for _, i := range agents {
			d.StreamListItem(ctx, i)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, nil
}

func getAgent(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_agent.getAgent", "connection_error", err)
		return nil, err
	}

	org := d.EqualsQuals["organization_slug"].GetStringValue()
	id := d.EqualsQuals["id"].GetStringValue()

	agent, resp, err := conn.Agents.Get(org, id)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_agent.getAgent", "query_error", err, "resp", resp, "org", org, "id", id)
		return nil, err
	}

	return agent, nil
}
