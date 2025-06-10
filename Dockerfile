FROM golang:1.23.2-alpine AS backend-builder
WORKDIR /app
COPY . .
RUN go mod download
RUN mkdir -p /app/auth /app/notifier /app/api-gateway /app/analytics /app/project
RUN  GOOS=linux go build -o /app/auth/auth ./services/auth/cmd/app/main.go
RUN  GOOS=linux go build -o /app/notifier/notifier ./services/notifier/cmd/app/main.go
RUN  GOOS=linux go build -o /app/api-gateway/api-gateway ./services/api-gateway/cmd/app/main.go
RUN  GOOS=linux go build -o /app/analytics/analytics ./services/analytics/cmd/app/main.go
RUN  GOOS=linux go build -o /app/project/project ./services/project/cmd/app/main.go
# Install wkhtmltopdf and its dependencies
COPY services/auth/.env /app/auth/.env
COPY services/auth/migrations /app/auth/migrations
COPY services/notifier/.env /app/notifier/.env
COPY services/api-gateway/.env /app/api-gateway/.env
COPY services/analytics/.env /app/analytics/.env
COPY services/project/.env /app/project/.env
COPY services/project/migrations /app/project/migrations
COPY services/analytics/migrations /app/analytics/migrations

FROM ghcr.io/surnet/alpine-wkhtmltopdf:3.21.3-0.12.6-full AS wkhtmltopdf

FROM golang:1.23.2-alpine AS auth-service
WORKDIR /app
RUN go install github.com/pressly/goose/v3/cmd/goose@v3.14.0
COPY --from=backend-builder /app/auth/auth .
COPY --from=backend-builder /app/auth/migrations /app/migrations
COPY --from=backend-builder /app/auth/.env .
RUN chmod +x /app/auth
EXPOSE 50051
CMD ["./auth", "--path", ".env"]

FROM golang:1.23.2-alpine AS analytics-service
WORKDIR /app
RUN mkdir -p /app/bin
RUN apk add --no-cache \
    libstdc++ \
    libx11 \
    libxrender \
    libxext \
    libssl3 \
    ca-certificates \
    fontconfig \
    freetype \
    ttf-dejavu \
    ttf-droid \
    ttf-freefont \
    && rm -rf /var/cache/apk/*

COPY --from=wkhtmltopdf /bin/wkhtmltopdf /usr/local/bin/wkhtmltopdf
COPY --from=backend-builder /app/analytics/analytics /app/bin/
COPY --from=backend-builder /app/analytics/.env .
COPY --from=backend-builder /app/analytics/migrations /app/migrations
RUN chmod +x /app/bin/analytics
RUN wkhtmltopdf --version
EXPOSE 50054
CMD ["/app/bin/analytics", "--path", ".env"]

FROM golang:1.23.2-alpine AS api-gateway
WORKDIR /app
COPY --from=backend-builder /app/api-gateway/api-gateway .
COPY --from=backend-builder /app/api-gateway/.env .
RUN chmod +x /app/api-gateway
EXPOSE 8080
CMD ["./api-gateway", "--path", ".env"]

FROM golang:1.23.2-alpine AS project-service
WORKDIR /app
COPY --from=backend-builder /app/project/project .
COPY --from=backend-builder /app/project/.env .
COPY --from=backend-builder /app/project/migrations /app/migrations
RUN chmod +x /app/project
EXPOSE 50053
CMD ["./project", "--path", ".env"]

FROM golang:1.23.2-alpine AS notifier-service
WORKDIR /app
COPY --from=backend-builder /app/notifier/notifier .
COPY --from=backend-builder /app/notifier/.env .
RUN chmod +x /app/notifier
EXPOSE 50052
CMD ["./notifier", "--path", ".env"]


