# Projects Service

## Migrations

```bash
goose -dir ./services/project/migrations sqlite "./services/project/data/pm.db" up
```

```bash
goose -dir ./services/project/migrations sqlite "./services/project/data/pm.db" down
```
