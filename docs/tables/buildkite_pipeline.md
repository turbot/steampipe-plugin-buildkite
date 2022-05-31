# Table: buildkite_pipeline

Pipelines defined in Buildkite.

## Examples

### List pipelines

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
