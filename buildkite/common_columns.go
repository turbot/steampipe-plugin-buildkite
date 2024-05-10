package buildkite

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/memoize"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func commonColumns(c []*plugin.Column) []*plugin.Column {
	return append([]*plugin.Column{
		{
			Name:        "member_id",
			Description: "Unique identifier for the organization.",
			Type:        proto.ColumnType_STRING,
			Hydrate:     getMemberId,
			Transform:   transform.FromValue(),
		},
	}, c...)
}

// if the caching is required other than per connection, build a cache key for the call and use it in Memoize.
var getMemberIdMemoized = plugin.HydrateFunc(getMemberIdUncached).Memoize(memoize.WithCacheKeyFunction(getMemberIdCacheKey))

// declare a wrapper hydrate function to call the memoized function
// - this is required when a memoized function is used for a column definition
func getMemberId(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	return getMemberIdMemoized(ctx, d, h)
}

// Build a cache key for the call to getMemberIdCacheKey.
func getMemberIdCacheKey(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {
	key := "getMemberId"
	return key, nil
}

func getMemberIdUncached(ctx context.Context, d *plugin.QueryData, h *plugin.HydrateData) (interface{}, error) {

	conn, err := connect(ctx, d)
	if err != nil {
		plugin.Logger(ctx).Error("getMemberIdUncached", "connection_error", err)
		return nil, err
	}

	user, resp, err := conn.User.Get()
	if err != nil {
		plugin.Logger(ctx).Error("getMemberIdUncached", "query_error", err, "resp", resp)
		return nil, err
	}

	return user.ID, nil
}
