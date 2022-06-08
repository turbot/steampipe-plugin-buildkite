![image](https://hub.steampipe.io/images/plugins/turbot/buildkite-social-graphic.png)

# Buildkite Plugin for Steampipe

Use SQL to query pipelines, builds, users and more from [Buildkite](https://buildkite.com).

* **[Get started →](https://hub.steampipe.io/plugins/turbot/buildkite)**
* Documentation: [Table definitions & examples](https://hub.steampipe.io/plugins/turbot/buildkite/tables)
* Community: [Slack Channel](https://steampipe.io/community/join)
* Get involved: [Issues](https://github.com/turbot/steampipe-plugin-buildkite/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io):

```shell
steampipe plugin install buildkite
```

Run steampipe:

```shell
steampipe query
```

Run a query:
```sql
  slug,
  name,
  id
from
  buildkite_organization
order by
  name;
```

```
+---------+------+--------------------------------------+
| slug    | name | id                                   |
+---------+------+--------------------------------------+
| test-74 | test | 11c22365-22ac-4368-aa11-ae1121919123 |
+---------+------+--------------------------------------+
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/turbot/steampipe-plugin-buildkite.git
cd steampipe-plugin-buildkite
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```
make
```

Configure the plugin:

```
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/buildkite.spc
```

Try it!

```
steampipe query
> .inspect buildkite
```

Further reading:
* [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
* [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Contributing

Please see the [contribution guidelines](https://github.com/turbot/steampipe/blob/main/CONTRIBUTING.md) and our [code of conduct](https://github.com/turbot/steampipe/blob/main/CODE_OF_CONDUCT.md). All contributions are subject to the [Apache 2.0 open source license](https://github.com/turbot/steampipe-plugin-prometheus/blob/main/LICENSE).

`help wanted` issues:
- [Steampipe](https://github.com/turbot/steampipe/labels/help%20wanted)
- [Buildkite Plugin](https://github.com/turbot/steampipe-plugin-buildkite/labels/help%20wanted)
