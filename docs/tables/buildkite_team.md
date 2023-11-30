---
title: "Steampipe Table: buildkite_team - Query Buildkite Teams using SQL"
description: "Allows users to query Teams in Buildkite, specifically the details of each team, providing insights into team members, permissions, and associated pipelines."
---

# Table: buildkite_team - Query Buildkite Teams using SQL

Buildkite Teams is a feature within Buildkite that allows for the grouping of users and pipelines. It provides a way to manage access control and permissions for pipelines based on teams rather than individual users. Buildkite Teams helps you maintain a structured and organized setup, making it easier to manage permissions and access to pipelines.

## Table Usage Guide

The `buildkite_team` table provides insights into Teams within Buildkite. As a DevOps engineer, explore team-specific details through this table, including members, permissions, and associated pipelines. Utilize it to uncover information about teams, such as members with specific permissions, the association between teams and pipelines, and the verification of access controls.

## Examples

### List all teams
Explore the organization's teams in Buildkite to understand the structure and hierarchy. This can be beneficial for administration and management, offering a clear view of the teams' arrangement.

```sql
select
  slug,
  name,
  id
from
  buildkite_team
order by
  name
```

### Teams created in the last week
Explore which teams have been recently established within the past week. This can be particularly useful for administrators to stay updated on new team formations and their details.

```sql
select
  slug,
  name,
  id,
  created_at
from
  buildkite_team
where
  created_at > now() - interval '7 days'
order by
  created_at desc
```

### Secret teams
Discover the segments that are classified as 'secret' within your Buildkite teams to enhance your understanding of privacy settings and ensure appropriate access control measures are in place.

```sql
select
  slug,
  name,
  id,
  privacy
from
  buildkite_team
where
  privacy = 'secret'
order by
  name
```