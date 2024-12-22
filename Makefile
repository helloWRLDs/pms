# Set environment
include .env

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
protoc:
	@echo "Generating protobuf files..."
	find ./services -name "*.proto" -print0 | xargs -0 -I {} \
	protoc -I ./services --go_out=./pkg/protobuf --go-grpc_out=./pkg/protobuf {}


# Run services
run:
	@echo "Running service: $(service)"
	go run ./pms.$(service)/cmd/app --path=./services/$(service)/.env