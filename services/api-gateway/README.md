# API Gateway Service

The API Gateway Service is the central entry point for all client requests in the PMS (Project Management System). It handles request routing, authentication, and service orchestration while providing a unified API interface for clients.

## Features

- **Request Routing**

  - Service discovery and routing
  - Load balancing
  - Request/response transformation
  - Protocol translation (gRPC to HTTP)

- **Authentication & Authorization**

  - JWT token validation
  - Session management
  - OAuth2 integration

- **Service Integration**
  - Auth service integration
  - Project service integration
  - Analytics service integration
  - Notification service integration

## Architecture

The service follows a modular architecture with the following components:

- **Router**: Request routing and middleware
- **Logic**: Business logic and service orchestration
- **Client**: Service client implementations
- **Models**: Data models and DTOs
- **Config**: Service configuration management

## Prerequisites

- Go 1.21 or higher
- Redis 7.0 or higher
- RabbitMQ 3.9 or higher
- Access to all microservices

## Configuration

The service can be configured using environment variables:

```env
# Service Configuration
HOST=:8080
FRONTEND_URL=http://localhost:3000

# JWT Configuration
JWT_SECRET=your-jwt-secret
JWT_TTL=24h

# Redis Configuration
REDIS_DSN=localhost:6379
REDIS_PASSWORD=
REDIS_DB=0

# Service Client Configurations
AUTH_HOST=localhost:50052
PROJECT_HOST=localhost:50051
ANALYTICS_HOST=localhost:50053

# Message Queue Configuration
NOTIFICATION_DSN=amqp://guest:guest@localhost:5672/
NOTIFICATION_EXCHANGE=notifications

# Logging Configuration
LOG_DEV=true
LOG_LEVEL=debug
LOG_FILE_OUTPUT=false
LOG_FILE_PATH=./logs/api-gateway.log
LOG_STACKTRACE=true
LOG_CALLER=true
```

## Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Configure environment variables (see Configuration section)
4. Run the service:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Authentication

- `POST /auth/login` - User login
- `POST /auth/register` - User registration
- `POST /auth/refresh` - Refresh access token
- `POST /auth/logout` - User logout

### Project Management

- `GET /projects` - List projects
- `POST /projects` - Create project
- `GET /projects/{id}` - Get project details
- `PUT /projects/{id}` - Update project
- `DELETE /projects/{id}` - Delete project

### Task Management

- `GET /tasks` - List tasks
- `POST /tasks` - Create task
- `GET /tasks/{id}` - Get task details
- `PUT /tasks/{id}` - Update task
- `DELETE /tasks/{id}` - Delete task

### Analytics

- `GET /analytics/projects` - Project analytics
- `GET /analytics/tasks` - Task analytics
- `GET /analytics/users` - User analytics

## Development

### Project Structure

```
internal/
├── client/     # Service clients
├── config/     # Configuration
├── logic/      # Business logic
├── models/     # Data models
└── router/     # Request routing
```

### Running Tests

```bash
go test ./...
```

### Local Development

1. Start Redis:

   ```bash
   docker run -d --name redis -p 6379:6379 redis:7
   ```

2. Start RabbitMQ:

   ```bash
   docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
   ```

3. Start the service in development mode:
   ```bash
   go run cmd/main.go
   ```

## Security Considerations

- JWT token validation
- Rate limiting per client
- Request size limits
- CORS configuration
- SSL/TLS encryption
- Input validation
- Error handling

## Monitoring

The service provides the following monitoring capabilities:

- Request/response metrics
- Error rates and types
- Service health checks
