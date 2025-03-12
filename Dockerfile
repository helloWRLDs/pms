FROM golang:1.23.2-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download
RUN mkdir -p /app/auth && GOOS=linux go build -o /app/auth/auth ./services/auth/cmd/app/main.go
RUN mkdir -p /app/notifier && GOOS=linux go build -o /app/notifier/notifier ./services/notifier/cmd/app/main.go
RUN mkdir -p /app/api-gateway && GOOS=linux go build -o /app/api-gateway/api-gateway ./services/api-gateway/cmd/app/main.go
COPY services/auth/.env /app/auth/.env
COPY services/auth/migrations /app/auth/migrations
COPY services/notifier/.env /app/notifier/.env
COPY services/api-gateway/.env /app/api-gateway/.env

FROM golang:1.23.2-alpine AS auth-service
WORKDIR /app
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY --from=builder /app/auth/auth .
COPY --from=builder /app/auth/migrations /app/migrations
COPY --from=builder /app/auth/.env .
RUN chmod +x /app/auth
EXPOSE 50051
CMD ["./auth", "--path", ".env"]
