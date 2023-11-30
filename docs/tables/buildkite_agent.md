---
title: "Steampipe Table: buildkite_agent - Query Buildkite Agents using SQL"
description: "Allows users to query Buildkite Agents, specifically providing insights into the status, version, and configuration of each agent."
---

# Table: buildkite_agent - Query Buildkite Agents using SQL

Buildkite Agents are small, reliable, and powerful execution units that run Buildkite jobs. They are responsible for running the actual build jobs, reporting the log output, and capturing the exit status code of the command that is run. Buildkite Agents can run on multiple platforms including Linux, macOS, Windows, FreeBSD, and Docker.

## Table Usage Guide

The `buildkite_agent` table provides insights into the individual agents within the Buildkite platform. As a DevOps engineer, explore agent-specific details through this table, including their status, version, and configuration. Utilize it to monitor the performance and health of your Buildkite Agents, and to ensure they are configured correctly and running the latest versions.

## Examples

### Most recent 3 agents registered
Explore the most recently registered agents within your system. This can help you to quickly identify new agents and ensure they are set up correctly.

```sql
select
  id,
  name,
  hostname,
  ip_address,
  created_at
from
  buildkite_agent
order by
  created_at desc
limit 3
```

### Agents by connection state
Analyze the distribution of Buildkite agents based on their connection states to better manage resources and troubleshoot connection issues. This can help optimize agent usage and ensure seamless operation of the Buildkite platform.

```sql
select
  connection_state,
  count(*)
from
  buildkite_agent
group by
  connection_state
order by
  count desc
```

### Agents by org
Analyze the distribution of Buildkite agents across various organizations. This can help in understanding which organizations have the most agents, aiding in resource allocation and management strategies.

```sql
select
  organization_slug,
  count(*)
from
  buildkite_agent
group by
  organization_slug
order by
  count desc
```

### Agents that haven't run any jobs in the last 24 hours
Explore which agents have been inactive in the last 24 hours. This can be useful to identify any underutilized resources and optimize your system's efficiency.

```sql
select
  id,
  name,
  hostname,
  ip_address,
  last_job_finished_at
from
  buildkite_agent
where
  last_job_finished_at < now() - interval '24 hours'
order by
  name
```