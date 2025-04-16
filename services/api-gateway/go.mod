module pms.api-gateway

go 1.23.2

require (
	github.com/gofiber/fiber/v2 v2.52.6
	github.com/sirupsen/logrus v1.9.3
	go.uber.org/zap v1.27.0
	pms.pkg v0.0.0
)

require (
	github.com/fasthttp/websocket v1.5.8 // indirect
	github.com/golang-jwt/jwt/v5 v5.2.1 // indirect
	github.com/redis/go-redis/v9 v9.7.0 // indirect
	github.com/savsgio/gotils v0.0.0-20240303185622-093b76447511 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)

replace pms.pkg => ../../pkg

require (
	github.com/andybalholm/brotli v1.1.0 // indirect
	github.com/caarlos0/env/v11 v11.3.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/gofiber/contrib/websocket v1.3.4
	github.com/google/uuid v1.6.0 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
	github.com/klauspost/compress v1.17.9 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.16 // indirect
	github.com/rivo/uniseg v0.2.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.52.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/net v0.34.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250115164207-1a7da9e5054f // indirect
	google.golang.org/grpc v1.71.0
	google.golang.org/protobuf v1.36.4 // indirect
)
