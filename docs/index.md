---
organization: Turbot
category: ["network"]
icon_url: "/images/plugins/turbot/buildkite.svg"
brand_color: "#14CC80"
display_name: Buildkite
name: buildkite
description: Steampipe plugin to query Buildkite pipelines, builds, users and more.
og_description: Query Buildkite with SQL! Open source CLI. No DB required.
og_image: "/images/plugins/turbot/buildkite-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Buildkite + Steampipe

[Steampipe](https://steampipe.io) is an open source CLI to instantly query cloud APIs using SQL.

[Buildkite](https://buildkite.com) is a platform for running fast, secure, and scalable continuous integration pipelines on your own infrastructure.

Example query:
```sql
select
  number,
  state,
  branch,
  blocked
from
  buildkite_build
where
  blocked
  and created_at > now() - interval '1 day';
```

```
+--------+-----------+--------+---------+
| number | state     | branch | blocked |
+--------+-----------+--------+---------+
| 1      | scheduled | master | true    |
+--------+-----------+--------+---------+
```

## Documentation

- **[Table definitions & examples →](/plugins/turbot/buildkite/tables)**

## Get started

### Install

Download and install the latest Buildkite plugin:

```bash
steampipe plugin install buildkite
```

### Configuration

Installing the latest buildkite plugin will create a config file (`~/.steampipe/config/buildkite.spc`) with a single connection named `buildkite`:

```hcl
connection "buildkite" {
  plugin = "buildkite"
  token  = "f7c8ce159f8e65f8f9abf0e655aa1b1afa5cef0c"
}
```

* `token` - Buildkite [API token](https://buildkite.com/docs/apis/managing-api-tokens).

Environment variables are also available as an alternate configuration method:
* `BUILDKITE_TOKEN`

### Permissions

The `token` should be assigned read permissions:
- `read_agents`
- `read_artifacts`
- `read_build_logs`
- `read_builds`
- `read_job_env`
- `read_organizations`
- `read_pipelines`
- `read_teams`
- `read_user`

## Get involved

* Open source: https://github.com/turbot/steampipe-plugin-buildkite
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
