FROM golang:1.23.2-alpine as builder
WORKDIR /app
COPY . .
RUN go mod download

RUN mkdir -p /app/auth /app/notifier /app/api-gateway /app/analytics /app/project

RUN  GOOS=linux go build -o /app/auth/auth ./services/auth/cmd/app/main.go

RUN  GOOS=linux go build -o /app/notifier/notifier ./services/notifier/cmd/app/main.go

RUN  GOOS=linux go build -o /app/api-gateway/api-gateway ./services/api-gateway/cmd/app/main.go

RUN  GOOS=linux go build -o /app/analytics/analytics ./services/analytics/cmd/app/main.go

RUN  GOOS=linux go build -o /app/project/project ./services/project/cmd/app/main.go

ADD https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.6/wkhtmltox-0.12.6-1.alpine3.14_amd64.apk /wkhtmltox.apk

RUN apk add --no-cache /wkhtmltox.apk

COPY services/auth/.env /app/auth/.env

COPY services/auth/migrations /app/auth/migrations

COPY services/notifier/.env /app/notifier/.env

COPY services/api-gateway/.env /app/api-gateway/.env

COPY services/analytics/.env /app/analytics/.env

FROM golang:1.23.2-alpine AS auth-service
WORKDIR /app
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
COPY --from=builder /app/auth/auth .
COPY --from=builder /app/auth/migrations /app/migrations
COPY --from=builder /app/auth/.env .
RUN chmod +x /app/auth
EXPOSE 50051
CMD ["./auth", "--path", ".env"]

FROM golang:1.23.2-alpine AS analytics-service
WORKDIR /app
COPY --from=builder /app/analytics/analytics .
COPY --from=builder /app/analytics/.env .
RUN chmod +x /app/analytics
EXPOSE 50054
CMD ["./analytics", "--path", ".env"]

FROM golang:1.23.2-alpine AS api-gateway
WORKDIR /app
COPY --from=builder /app/api-gateway/api-gateway .
COPY --from=builder /app/api-gateway/.env .
RUN chmod +x /app/api-gateway
EXPOSE 50053
CMD ["./api-gateway", "--path", ".env"]

FROM golang:1.23.2-alpine AS project-service
WORKDIR /app
COPY --from=builder /app/project/project .
COPY --from=builder /app/project/.env .
RUN chmod +x /app/project
EXPOSE 50053
CMD ["./project", "--path", ".env"]

FROM golang:1.23.2-alpine AS notifier-service
WORKDIR /app
COPY --from=builder /app/notifier/notifier .
COPY --from=builder /app/notifier/.env .
RUN chmod +x /app/notifier
EXPOSE 50052
CMD ["./notifier", "--path", ".env"]
