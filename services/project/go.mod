module pms.project

go 1.23.2

require (
	github.com/jmoiron/sqlx v1.4.0
	go.uber.org/zap v1.27.0
	google.golang.org/protobuf v1.36.3
	pms.pkg v0.0.0
)

replace pms.pkg => ../../pkg

require (
	github.com/caarlos0/env/v11 v11.3.1 // indirect
	github.com/cznic/ccgo v0.0.0-20181122101908-2735556262de // indirect
	github.com/cznic/ccir v0.0.0-20181122101719-1a0aa558e495 // indirect
	github.com/cznic/internal v0.0.0-20181122101858-3279554c546e // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/cznic/memory v0.0.0-20181122101858-44f9dcde99e8 // indirect
	github.com/cznic/sqlite v0.0.0-20181122101901-0b3d034df24f // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.69.4 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/Masterminds/squirrel v1.5.4
	github.com/google/uuid v1.6.0
	github.com/stretchr/testify v1.10.0
	go.uber.org/multierr v1.10.0 // indirect
)
