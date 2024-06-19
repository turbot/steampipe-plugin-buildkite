package buildkite

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableBuildkiteUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "buildkite_user",
		Description: "User in the Buildkite account.",
		List: &plugin.ListConfig{
			Hydrate: listUser,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Description: "ID of the user."},
			{Name: "name", Type: proto.ColumnType_STRING, Description: "Name of the user."},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the user."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("CreatedAt").Transform(timeToRfc3339), Description: "Time when the user was created."},
			{Name: "access_token_uuid", Type: proto.ColumnType_STRING, Hydrate: getAccessToken, Transform: transform.FromField("UUID"), Description: "UUID of the access token for these credentials."},
			{Name: "scopes", Type: proto.ColumnType_JSON, Hydrate: getAccessToken, Description: "Scopes assigned to the access token for this user."},
		}),
	}
}

func listUser(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_user.listUser", "connection_error", err)
		return nil, err
	}

	user, resp, err := conn.User.Get()
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_user.listUser", "query_error", err, "resp", resp)
		return nil, err
	}

	d.StreamListItem(ctx, user)

	return nil, nil
}

func getAccessToken(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_user.listUser", "connection_error", err)
		return nil, err
	}

	at, resp, err := conn.AccessTokens.Get()
	if err != nil {
		plugin.Logger(ctx).Error("buildkite_user.getAccessToken", "query_error", err, "resp", resp)
		return nil, err
	}

	return at, nil
}
