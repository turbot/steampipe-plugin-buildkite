# Table: buildkite_user

Information about the user configured for this Buildkite connection.

## Examples

### Get user info

```sql
select
  *
from
  buildkite_user
```

### Scopes assigned to the access token for this user

```sql
select
  jsonb_array_elements_text(scopes) as scope
from
  buildkite_user
```
