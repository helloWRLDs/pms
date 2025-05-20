module pms.auth

go 1.23.2

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/google/uuid v1.6.0
	github.com/jmoiron/sqlx v1.4.0
	github.com/lib/pq v1.10.9
	github.com/stretchr/testify v1.10.0
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.36.0
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.6
	pms.pkg v0.0.0
)

replace pms.pkg => ../../pkg

require (
	github.com/caarlos0/env/v11 v11.3.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-sql-driver/mysql v1.9.1 // indirect
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	go.opentelemetry.io/otel v1.35.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	golang.org/x/sys v0.31.0 // indirect
)
