---
title: "Steampipe Table: buildkite_pipeline - Query Buildkite Pipelines using SQL"
description: "Allows users to query Buildkite Pipelines, providing insights into pipeline configurations, steps, and related metadata."
---

# Table: buildkite_pipeline - Query Buildkite Pipelines using SQL

Buildkite Pipelines are a key component of the Buildkite CI/CD platform. They define a series of steps that are executed in order to test, build, and deploy code. Each pipeline is associated with a specific code repository, and can be configured to trigger on various events, such as code commits or pull request updates.

## Table Usage Guide

The `buildkite_pipeline` table provides detailed information about each pipeline in your Buildkite account. Developers, testers, and DevOps engineers can use this table to examine pipeline configurations, identify the steps involved in each pipeline, and analyze the status of code builds and deployments. This can be particularly useful for troubleshooting build failures, optimizing build times, and ensuring consistent deployment practices.

## Examples

### List pipelines
Explore the Buildkite pipelines in your system, ordered by their names, to gain insights into the different processes running in your environment. This can be useful for managing and optimizing your workflows.

```sql
select
  slug,
  name
from
  buildkite_pipeline
order by
  name
```

### Pipelines with waiting jobs
Explore which pipelines have jobs pending execution to prioritize resources and manage workflow efficiently. This allows for better allocation of resources and minimizes downtime in the pipeline.

```sql
select
  slug,
  name,
  waiting_jobs_count
from
  buildkite_pipeline
where
  waiting_jobs_count > 0
order by
  waiting_jobs_count desc
```

### Pipelines with the ENV VAR AWS_ACCESS_KEY_ID set
Explore which pipelines have the 'AWS_ACCESS_KEY_ID' environment variable set. This is useful to identify potential security risks or misconfigurations in your Buildkite pipelines.

```sql
select
  slug,
  name,
  env
from
  buildkite_pipeline
where
  env ? 'AWS_ACCESS_KEY_ID'
```