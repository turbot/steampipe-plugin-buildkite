package buildkite

import (
	"context"

	bkapi "github.com/buildkite/go-buildkite/v3/buildkite"

	"github.com/turbot/steampipe-plugin-sdk/v4/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func tableBuildkitePipeline(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "buildkite_pipeline",
		Description: "Pipelines in the Buildkite account.",
		List: &plugin.ListConfig{
			ParentHydrate: listOrganization,
			Hydrate:       listPipeline,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "organization_slug", Require: plugin.Optional},
			},
		},
		Get: &plugin.GetConfig{
			Hydrate: getPipeline,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "organization_slug"},
				{Name: "slug"},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "slug", Type: proto.ColumnType_STRING, Description: "Slug of the pipeline."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the pipeline."},
			{Name: "organization_slug", Type: proto.ColumnType_STRING, Transform: transform.FromField("URL").Transform(orgFromURL), Description: "Organization of the pipeline."},
			// Other columns
			{Name: "archived_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("ArchivedAt").Transform(timeToRfc3339), Description: "Time when the pipeline was archived."},
			{Name: "badge_url", Type: proto.ColumnType_STRING, Description: "Badge URL for the pipeline."},
			{Name: "branch_configuration", Type: proto.ColumnType_STRING, Description: "A branch filter pattern to limit which pushed branches trigger builds on this pipeline."},
			{Name: "builds_url", Type: proto.ColumnType_STRING, Description: "Builds URL for the pipeline."},
			{Name: "cancel_running_branch_builds", Type: proto.ColumnType_BOOL, Description: "If true then when a new build is created on a branch, any previous builds that are running on the same branch will be automatically canceled."},
			{Name: "cancel_running_branch_builds_filter", Type: proto.ColumnType_STRING, Description: "A branch filter pattern to limit which branches intermediate build cancelling applies to."},
			{Name: "cluster_id", Type: proto.ColumnType_STRING, Description: "ID of the cluster for this pipeline."},
			{Name: "configuration", Type: proto.ColumnType_STRING, Description: "The YAML pipeline that consists of the build pipeline steps."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timeToRfc3339), Description: "Time when the pipeline was created."},
			{Name: "default_branch", Type: proto.ColumnType_STRING, Description: "The name of the branch to prefill when new builds are created or triggered in Buildkite."},
			{Name: "description", Type: proto.ColumnType_STRING, Description: "Description of the pipeline."},
			{Name: "env", Type: proto.ColumnType_JSON, Description: "Environment variables for the pipeline."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the pipeline."},
			{Name: "provider", Type: proto.ColumnType_JSON, Description: "Provider information."},
			{Name: "repository", Type: proto.ColumnType_STRING, Description: "Repository for the pipeline."},
			{Name: "running_builds_count", Type: proto.ColumnType_INT, Transform: transform.FromField("RunningBuildsCount"), Description: "Number of running builds."},
			{Name: "running_jobs_count", Type: proto.ColumnType_INT, Transform: transform.FromField("RunningBuildsCount"), Description: "Number of running jobs."},
			{Name: "scheduled_builds_count", Type: proto.ColumnType_INT, Transform: transform.FromField("ScheduledBuildsCount"), Description: "Number of scheduled builds."},
			{Name: "scheduled_jobs_count", Type: proto.ColumnType_INT, Transform: transform.FromField("ScheduledJobsCount"), Description: "Number of scheduled jobs."},
			{Name: "skip_queued_branch_builds", Type: proto.ColumnType_BOOL, Description: "If true then when a new build is created on a branch, any previous builds that haven't yet started on the same branch will be automatically marked as skipped."},
			{Name: "skip_queued_branch_builds_filter", Type: proto.ColumnType_STRING, Description: "A branch filter pattern to limit which branches intermediate build skipping applies to."},
			{Name: "steps", Type: proto.ColumnType_JSON, Description: "Build step definitions."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL for the pipeline."},
			{Name: "visibility", Type: proto.ColumnType_STRING, Description: "Visibility of the pipeline."},
			{Name: "waiting_jobs_count", Type: proto.ColumnType_INT, Transform: transform.FromField("WaitingJobsCount"), Description: "Number of waiting jobs."},
			{Name: "web_url", Type: proto.ColumnType_STRING, Description: "Web URL for the pipeline."},
		},
	}
}

func listPipeline(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_pipeline.listPipeline", "connection_error", err)
		return nil, err
	}

	var org string
	if h.Item != nil {
		org = *h.Item.(bkapi.Organization).Slug
	} else if d.KeyColumnQuals["organization_slug"] != nil {
		org = d.KeyColumnQuals["organization_slug"].GetStringValue()
	}

	opts := &bkapi.PipelineListOptions{
		ListOptions: bkapi.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}
	for {
		pipelines, resp, err := conn.Pipelines.List(org, opts)
		if err != nil {
			plugin.Logger(ctx).Error("buildkite_pipeline.listPipeline", "query_error", err, "resp", resp)
			return nil, err
		}
		for _, i := range pipelines {
			d.StreamListItem(ctx, i)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, nil
}

func getPipeline(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_pipeline.getPipeline", "connection_error", err)
		return nil, err
	}

	org := d.KeyColumnQuals["organization_slug"].GetStringValue()
	slug := d.KeyColumnQuals["slug"].GetStringValue()

	pipeline, resp, err := conn.Pipelines.Get(org, slug)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_pipeline.getPipeline", "query_error", err, "resp", resp, "org", org, "slug", slug)
		return nil, err
	}

	return pipeline, nil
}
