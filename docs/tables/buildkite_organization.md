# Table: buildkite_organization

Organizations the user has access to query.

## Examples

### List all organizations

```sql
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

```sql
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
