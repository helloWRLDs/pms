# Set environment
# include .env

path:
	@export GOPATH=$$HOME/go
	@export PATH=$$PATH:$$GOPATH/bin

# installers
grpc.init: path
	sudo apt-get update && sudo apt-get install -y protobuf-compiler
	@echo "Installing protoc-gen-go, protoc-gen-go-grpc"
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# Generate gRPC code from .proto files
genproto:
	@echo "Generating protobuf files..."
	protoc --go_out=. --go-grpc_out=. .\pkg\transport\grpc\proto\${name}.proto

protoc:
	@echo "Generating protobuf files..."
	protoc  --go_out=. --go-grpc_out=. .\services\${service}\proto\${service}_service.proto
	protoc  --go_out=. --go-grpc_out=. .\services\${service}\proto\${service}.proto

# Run services
run:
	@echo "Running service: $(service)"
	go run ./services/$(service)/cmd/app --path=./services/$(service)/.env

migrate.up:
	goose -dir ./services/analytics/migrations postgres "postgres://postgres:postgres@localhost:5432/analytics?sslmode=disable" up
	goose -dir ./services/project/migrations postgres "postgres://postgres:postgres@localhost:5432/project?sslmode=disable" up
	goose -dir ./services/auth/migrations postgres "postgres://postgres:postgres@localhost:5432/auth?sslmode=disable" up