# Authorization Service

## ERD

![erd](./docs/auth_erd.png)

## Migrations

```bash
goose -dir ./migrations sqlite "./data/users.db" up
```

```bash
goose -dir ./migrations sqlite "./data/users.db" down
```
