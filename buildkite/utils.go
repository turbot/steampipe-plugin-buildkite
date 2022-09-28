package buildkite

import (
	"context"
	"errors"
	"os"
	"strings"
	"time"

	bkapi "github.com/buildkite/go-buildkite/v3/buildkite"

	"github.com/turbot/go-kit/types"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v4/plugin/transform"
)

func connect(ctx context.Context, d *plugin.QueryData) (*bkapi.Client, error) {

	// Load connection from cache, which preserves throttling protection etc
	cacheKey := "buildkite"
	if cachedData, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		return cachedData.(*bkapi.Client), nil
	}

	// Default to using env vars
	token := os.Getenv("BUILDKITE_TOKEN")

	// But prefer the config
	buildkiteConfig := GetConfig(d.Connection)
	if buildkiteConfig.Token != nil {
		token = *buildkiteConfig.Token
	}

	if token == "" {
		// Credentials not set
		return nil, errors.New("token must be configured")
	}

	config, err := bkapi.NewTokenConfig(token, false)
	if err != nil {
		plugin.Logger(ctx).Error("buildkite.connect", "config_error", err)
		return nil, err
	}

	client := bkapi.NewClient(config.Client())

	// Save to cache
	d.ConnectionManager.Cache.Set(cacheKey, client)

	return client, nil
}

func isNotFoundError(err error) bool {
	return strings.Contains(err.Error(), "404 Not Found")
}

func timeToRfc3339(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value != nil {
		t := d.Value.(*bkapi.Timestamp)
		if t != nil {
			return t.Format(time.RFC3339), nil
		}
	}
	return nil, nil
}

func orgFromURL(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	if d.Value == nil {
		return nil, nil
	}
	u := types.SafeString(d.Value)
	parts := strings.Split(u, "/")
	return parts[5], nil
}
