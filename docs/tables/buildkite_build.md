---
title: "Steampipe Table: buildkite_build - Query Buildkite Builds using SQL"
description: "Allows users to query Buildkite Builds, providing detailed information about each build, including its state, number, commit, branch, message, and more."
---

# Table: buildkite_build - Query Buildkite Builds using SQL

Buildkite is a continuous integration and deployment tool that integrates with version control systems to run tests on your code. It is designed to work with your existing tools and workflows, and allows you to define your build pipelines in code. Buildkite provides a platform for running fast, secure, and scalable pipelines on your own infrastructure.

## Table Usage Guide

The `buildkite_build` table provides insights into Builds within Buildkite. As a DevOps engineer, you can explore build-specific details through this table, including build state, number, commit, branch, and more. Utilize it to uncover information about each build, such as its current status, associated commit, and the branch it belongs to.

## Examples

### Builds created in the last 15 mins
Gain insights into recent activity by identifying builds that have been initiated in the last 15 minutes. This allows for immediate awareness and response to any new developments or issues.

```sql
select
  organization_slug,
  pipeline_slug,
  number,
  created_at,
  state
from
  buildkite_build
where
  created_at > now() - interval '15 mins'
order by
  created_at desc
```

### Builds by org
Determine the areas in which different organizations and pipelines are contributing the most by analyzing the frequency of builds. This can be useful for understanding resource allocation and identifying heavily utilized pipelines within specific organizations.

```sql
select
  organization_slug,
  pipeline_slug,
  count(*)
from
  buildkite_build
group by
  organization_slug,
  pipeline_slug
order by
  count desc
```

### Builds by day over the last week
Discover the frequency of builds started in the last two weeks. This query helps to understand the build activity patterns and could be useful for identifying peak times or days for build initiations.

```sql
select
  date_part('date', started_at) as hour,
  count(*)
from
  buildkite_build
where
  started_at > now() - interval '14 days'
group by
  hour
order by
  hour desc
```