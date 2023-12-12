---
title: "Steampipe Table: buildkite_organization - Query Buildkite Organizations using SQL"
description: "Allows users to query Buildkite Organizations, specifically to retrieve information about the organizations' settings, members, pipelines, and teams."
---

# Table: buildkite_organization - Query Buildkite Organizations using SQL

Buildkite is a platform that helps developers automate their software build and testing processes. The platform's organization resource represents a group of users that have access to a shared set of pipelines. Organizations in Buildkite can have multiple teams and pipelines, and their settings can be configured to meet specific requirements.

## Table Usage Guide

The `buildkite_organization` table provides insights into organizations within Buildkite. As a DevOps engineer, explore organization-specific details through this table, including settings, members, pipelines, and teams. Utilize it to uncover information about organizations, such as their members' access levels, the configuration of their pipelines, and the structure of their teams.

## Examples

### List all organizations
Explore the different organizations within your system, ordered by their names, to better manage and understand your organizational structure. This can be particularly useful for administrators or managers who need a comprehensive overview of all the organizations they oversee.

```sql+postgres
select
  slug,
  name,
  id
from
  buildkite_organization
order by
  name;
```

```sql+sqlite
select
  slug,
  name,
  id
from
  buildkite_organization
order by
  name;
```

### Organizations created in the last week
Explore which organizations have been established in the recent week. This can be useful for keeping track of new additions and assessing the growth rate of your network.

```sql+postgres
select
  slug,
  name,
  id,
  created_at
from
  buildkite_organization
where
  created_at > now() - interval '7 days'
order by
  created_at desc;
```

```sql+sqlite
select
  slug,
  name,
  id,
  created_at
from
  buildkite_organization
where
  created_at > datetime('now', '-7 days')
order by
  created_at desc;
```