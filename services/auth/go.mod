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
	github.com/cznic/ccgo v0.0.0-20181122101908-2735556262de // indirect
	github.com/cznic/ccir v0.0.0-20181122101719-1a0aa558e495 // indirect
	github.com/cznic/internal v0.0.0-20181122101858-3279554c546e // indirect
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/cznic/memory v0.0.0-20181122101858-44f9dcde99e8 // indirect
	github.com/cznic/sqlite v0.0.0-20181122101901-0b3d034df24f // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mfridman/interpolate v0.0.2 // indirect
	github.com/ncruces/go-strftime v0.1.9 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230129092748-24d4a6f8daec // indirect
	github.com/sethvargo/go-retry v0.3.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20250305212735-054e65f0b394 // indirect
	golang.org/x/net v0.38.0 // indirect
	golang.org/x/sync v0.12.0 // indirect
	golang.org/x/text v0.23.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250324211829-b45e905df463 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	modernc.org/libc v1.62.1 // indirect
	modernc.org/mathutil v1.7.1 // indirect
	modernc.org/memory v1.9.1 // indirect
)

require (
	github.com/golang-jwt/jwt/v5 v5.2.2
	github.com/lann/builder v0.0.0-20180802200727-47ae307949d0 // indirect
	github.com/lann/ps v0.0.0-20150810152359-62de8c46ede0 // indirect
	github.com/pressly/goose/v3 v3.24.2
	github.com/sirupsen/logrus v1.9.3
	golang.org/x/sys v0.31.0 // indirect
	modernc.org/sqlite v1.37.0
)
