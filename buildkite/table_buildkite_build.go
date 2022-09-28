package buildkite

import (
	"context"
	"time"

	bkapi "github.com/buildkite/go-buildkite/v3/buildkite"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableBuildkiteBuild(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "buildkite_build",
		Description: "Builds in the Buildkite account.",
		List: &plugin.ListConfig{
			Hydrate: listBuild,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "branch", Require: plugin.Optional},
				{Name: "commit", Require: plugin.Optional},
				{Name: "created_at", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional, CacheMatch: "exact"},
				{Name: "finished_at", Operators: []string{">", ">=", "=", "<", "<="}, Require: plugin.Optional, CacheMatch: "exact"},
				{Name: "state", Require: plugin.Optional},
			},
		},
		/*
			// Builds.Get fails in SDK with Error: json: cannot unmarshal array into Go value of type buildkite.Build (SQLSTATE HV000)
			Get: &plugin.GetConfig{
				Hydrate: getBuild,
				KeyColumns: []*plugin.KeyColumn{
					{Name: "organization_slug"},
					{Name: "pipeline_slug"},
					{Name: "number"},
				},
			},
		*/
		Columns: []*plugin.Column{
			// Top columns
			{Name: "organization_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL").Transform(orgFromURL), Description: "Organization of the build."},
			{Name: "pipeline_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("Pipeline.Slug"), Description: "Slug of the pipeline the build is for."},
			{Name: "number", Type: proto.ColumnType_INT, Description: "Number of the build."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the build."},
			// Other columns
			{Name: "blocked", Type: proto.ColumnType_BOOL, Description: "True if the build is blocked."},
			{Name: "branch", Type: proto.ColumnType_STRING, Description: "Branch used for the build."},
			{Name: "commit", Type: proto.ColumnType_STRING, Description: "Commit used for the build."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timeToRfc3339), Description: "Time when the build was created."},
			{Name: "creator", Type: proto.ColumnType_JSON, Description: "Creator of the build."},
			{Name: "env", Type: proto.ColumnType_JSON, Description: "Environment variables used for the build."},
			{Name: "finished_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("FinishedAt").Transform(timeToRfc3339), Description: "Time when the build finished."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the build."},
			{Name: "jobs", Type: proto.ColumnType_JSON, Description: "Jobs run during the build."},
			{Name: "message", Type: proto.ColumnType_STRING, Description: "Message for the build."},
			{Name: "meta_data", Type: proto.ColumnType_JSON, Description: "Metadata for the build."},
			{Name: "pipeline", Type: proto.ColumnType_JSON, Description: "Pipeline the build is for."},
			{Name: "pull_request", Type: proto.ColumnType_JSON, Description: "Pull request for the build."},
			{Name: "scheduled_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ScheduledAt").Transform(timeToRfc3339), Description: "Time when the build was scheduled."},
			{Name: "started_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("StartedAt").Transform(timeToRfc3339), Description: "Time when the build was started."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL for the build."},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "Web URL for the build."},
		},
	}
}

func listBuild(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_build.listBuild", "connection_error", err)
		return nil, err
	}

	opts := &bkapi.BuildsListOptions{
		ListOptions: bkapi.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}

	if d.KeyColumnQuals["branch"] != nil {
		opts.Branch = d.KeyColumnQuals["branch"].GetStringValue()
	}

	if d.KeyColumnQuals["commit"] != nil {
		opts.Commit = d.KeyColumnQuals["commit"].GetStringValue()
	}

	if d.KeyColumnQuals["state"] != nil {
		opts.State = []string{d.KeyColumnQuals["state"].GetStringValue()}
	}

	// Find the time range from optional quals.
	createdAtMinTime := time.Time{}
	createdAtMaxTime := time.Time{}
	if d.Quals["created_at"] != nil {
		for _, q := range d.Quals["created_at"].Quals {
			ts := q.Value.GetTimestampValue().AsTime()
			switch q.Operator {
			case ">=", ">":
				if ts.After(createdAtMinTime) {
					createdAtMinTime = ts
				}
			case "=":
				createdAtMinTime = ts
				createdAtMaxTime = ts
			case "<", "<=":
				if ts.Before(createdAtMaxTime) {
					createdAtMaxTime = ts
				}
			}
		}
	}
	if !createdAtMinTime.IsZero() {
		opts.CreatedFrom = createdAtMinTime
	}
	if !createdAtMaxTime.IsZero() {
		opts.CreatedTo = createdAtMaxTime
	}

	// Find the time range from optional quals.
	finishedAtMinTime := time.Time{}
	if d.Quals["finished_at"] != nil {
		for _, q := range d.Quals["finished_at"].Quals {
			ts := q.Value.GetTimestampValue().AsTime()
			switch q.Operator {
			case ">=", ">":
				{
					if ts.After(finishedAtMinTime) {
						finishedAtMinTime = ts
					}
				}
			case "=":
				{
					finishedAtMinTime = ts
					break
				}
			}
		}
	}
	if !finishedAtMinTime.IsZero() {
		opts.FinishedFrom = finishedAtMinTime
	}

	for {
		builds, resp, err := conn.Builds.List(opts)
		if err != nil {
			plugin.Logger(ctx).Error("buildkite_build.listBuild", "query_error", err, "resp", resp)
			return nil, err
		}
		for _, i := range builds {
			d.StreamListItem(ctx, i)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, nil
}

/*
// Builds.Get fails in SDK with Error: json: cannot unmarshal array into Go value of type buildkite.Build (SQLSTATE HV000)
func getBuild(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_build.getBuild", "connection_error", err)
		return nil, err
	}

	org := d.KeyColumnQuals["organization_slug"].GetStringValue()
	pipeline := d.KeyColumnQuals["pipeline_slug"].GetStringValue()
	num := d.KeyColumnQuals["number"].GetStringValue()

	build, resp, err := conn.Builds.Get(org, pipeline, num, nil)

	if err != nil {
		plugin.Logger(ctx).Error("buildkite_build.getBuild", "query_error", err, "resp", resp, "org", org, "pipeline", pipeline, "num", num)
		return nil, err
	}

	return build, nil
}
*/
