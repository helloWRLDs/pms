# Authorization Service

## ERD

![erd](./docs/auth_erd.png)

## Migrations

```bash
goose -dir ./services/auth/migrations sqlite "./services/auth/data/users.db" up
```

```bash
goose -dir ./services/auth/migrations sqlite "./services/auth/data/users.db" down
```
