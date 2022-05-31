# Table: buildkite_team

Teams in Buildkite.

## Examples

### List all teams

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
