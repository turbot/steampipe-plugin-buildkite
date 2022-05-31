# Table: buildkite_agent

Agents registered with Buildkite.

## Examples

### Most recent 3 agents registered

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
