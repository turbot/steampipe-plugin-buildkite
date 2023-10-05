## v0.3.1 [2023-10-05]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.6.2](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v562-2023-10-03) which prevents nil pointer reference errors for implicit hydrate configs. ([#21](https://github.com/turbot/steampipe-plugin-buildkite/pull/21))

## v0.3.0 [2023-10-02]

_Dependencies_

- Upgraded to [steampipe-plugin-sdk v5.6.1](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v561-2023-09-29) with support for rate limiters. ([#18](https://github.com/turbot/steampipe-plugin-buildkite/pull/18))
- Recompiled plugin with Go version `1.21`. ([#18](https://github.com/turbot/steampipe-plugin-buildkite/pull/18))

## v0.2.0 [2023-04-06]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v5.3.0](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v530-2023-03-16) which includes fixes for query cache pending item mechanism and aggregator connections not working for dynamic tables. ([#10](https://github.com/turbot/steampipe-plugin-buildkite/pull/10))

## v0.1.0 [2022-09-28]

_Dependencies_

- Recompiled plugin with [steampipe-plugin-sdk v4.1.7](https://github.com/turbot/steampipe-plugin-sdk/blob/main/CHANGELOG.md#v417-2022-09-08) which includes several caching and memory management improvements. ([#7](https://github.com/turbot/steampipe-plugin-buildkite/pull/7))
- Recompiled plugin with Go version `1.19`. ([#7](https://github.com/turbot/steampipe-plugin-buildkite/pull/7))

## v0.0.2 [2022-06-09]

_Bug fixes_

- Fixed the brand color of the plugin icon. ([#3](https://github.com/turbot/steampipe-plugin-buildkite/pull/3))

## v0.0.1 [2022-06-09]

_What's new?_

- New tables added
  - [buildkite_agent](https://hub.steampipe.io/plugins/turbot/buildkite/tables/buildkite_agent)
  - [buildkite_build](https://hub.steampipe.io/plugins/turbot/buildkite/tables/buildkite_build)
  - [buildkite_organization](https://hub.steampipe.io/plugins/turbot/buildkite/tables/buildkite_organization)
  - [buildkite_pipeline](https://hub.steampipe.io/plugins/turbot/buildkite/tables/buildkite_pipeline)
  - [buildkite_team](https://hub.steampipe.io/plugins/turbot/buildkite/tables/buildkite_team)
  - [buildkite_user](https://hub.steampipe.io/plugins/turbot/buildkite/tables/buildkite_user)
