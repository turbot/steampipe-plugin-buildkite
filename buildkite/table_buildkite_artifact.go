package buildkite

import (
	"context"
	"fmt"

	bkapi "github.com/buildkite/go-buildkite/v3/buildkite"

	"github.com/turbot/steampipe-plugin-sdk/v3/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v3/plugin/transform"
)

func tableBuildkiteArtifact(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "buildkite_artifact",
		Description: "Artifacts produced by builds in the Buildkite account.",
		List: &plugin.ListConfig{
			ParentHydrate: listBuild,
			Hydrate:       listArtifact,
			KeyColumns: []*plugin.KeyColumn{
				{Name: "build_number", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			// Top columns
			{Name: "build_number", Type: proto.ColumnType_INT, Transform: transform.FromQual("build_number"), Description: "Build number for the artifact."},
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the artifact."},
			{Name: "download_url", Type: proto.ColumnType_STRING, Description: "Download URL for the artifact."},
			// Other columns
			{Name: "dirname", Type: proto.ColumnType_STRING, Description: "Directory name of the artifact."},
			{Name: "file_size", Type: proto.ColumnType_INT, Description: "File size of the artifact."},
			{Name: "filename", Type: proto.ColumnType_STRING, Description: "File name of the artifact."},
			{Name: "glob_path", Type: proto.ColumnType_STRING, Description: "Glob path for the artifact."},
			{Name: "job_id", Type: proto.ColumnType_STRING, Description: "ID of the job for the artifact."},
			{Name: "mime_type", Type: proto.ColumnType_STRING, Description: "MIME type of the artifact."},
			{Name: "original_path", Type: proto.ColumnType_STRING, Description: "Original path for the artifact."},
			{Name: "path", Type: proto.ColumnType_STRING, Description: "Path of the artifact."},
			{Name: "sha1sum", Type: proto.ColumnType_STRING, Transform: transform.FromField("SHA1"), Description: "SHA1 sum of the artifact."},
			{Name: "state", Type: proto.ColumnType_STRING, Description: "State of the artifact."},
			{Name: "url", Type: proto.ColumnType_STRING, Description: "URL for the artifact."},
		},
	}
}

func listArtifact(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_artifact.listArtifact", "connection_error", err)
		return nil, err
	}

	var buildNumber, orgSlug, pipelineSlug string
	if h.Item != nil {
		build := h.Item.(bkapi.Build)
		buildNumber = fmt.Sprintf("%d", *build.Number)
		pipelineSlug = *build.Pipeline.Slug
		orgSlug = "turbot"
	} else if d.KeyColumnQuals["build_number"] != nil {
		buildNumber = d.KeyColumnQuals["build_number"].GetStringValue()
	}

	opts := &bkapi.ArtifactListOptions{
		ListOptions: bkapi.ListOptions{
			Page:    1,
			PerPage: 100,
		},
	}

	for {
		artifacts, resp, err := conn.Artifacts.ListByBuild(orgSlug, pipelineSlug, buildNumber, opts)
		if err != nil {
			plugin.Logger(ctx).Error("buildkite_artifact.listArtifact", "query_error", err, "resp", resp)
			return nil, err
		}
		for _, i := range artifacts {
			d.StreamListItem(ctx, i)
		}
		if resp.NextPage == 0 {
			break
		}
		opts.Page = resp.NextPage
	}

	return nil, nil
}
