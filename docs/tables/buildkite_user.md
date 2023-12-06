---
title: "Steampipe Table: buildkite_user - Query Buildkite Users using SQL"
description: "Allows users to query Buildkite Users, specifically providing insights into user profiles, including their email, name, and avatar URL."
---

# Table: buildkite_user - Query Buildkite Users using SQL

Buildkite is a platform that allows you to run fast, secure, and scalable continuous integration pipelines on your own infrastructure. It provides a way to manage and monitor your build pipelines, with features for parallel testing, real-time updates, and more. A Buildkite User is an individual with a registered account on the Buildkite platform, and can create, manage, or contribute to build pipelines.

## Table Usage Guide

The `buildkite_user` table provides insights into individual user profiles within the Buildkite platform. As a DevOps engineer, you can explore user-specific details through this table, including their email, name, and avatar URL. This can be useful for auditing user activity, managing user permissions, or investigating issues related to specific user accounts.

## Examples

### Get user info
Explore the details of all users in the Buildkite system to understand their roles and permissions. This can be useful for auditing purposes or to ensure the correct access levels are assigned to each user.

```sql+postgres
select
  *
from
  buildkite_user
```

```sql+sqlite
select
  *
from
  buildkite_user
```

### Scopes assigned to the access token for this user
Explore which scopes are assigned to the user's access token in Buildkite, allowing you to understand and manage the permissions and access levels of different users.

```sql+postgres
select
  jsonb_array_elements_text(scopes) as scope
from
  buildkite_user
```

```sql+sqlite
select
  json_each.value as scope
from
  buildkite_user,
  json_each(scopes)
```