# Table: buildkite_build

Builds across all your organizations.

## Examples

### Builds created in the last 15 mins

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
