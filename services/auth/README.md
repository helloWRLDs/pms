# Authorization Service

The Authorization Service is responsible for managing user authentication, authorization, and company management in the PMS (Project Management System). It provides secure user authentication through multiple providers and handles company-related operations.

## Features

- User Authentication
  - Email/Password authentication
  - Google OAuth2 integration
  - GitHub OAuth2 integration
- JWT-based session management
- Company Management
  - Company creation and management
  - Participant management
- User Management
  - User registration and profile management
  - Session management
  - Password reset functionality

## Architecture

### ERD

![erd](./docs/auth_erd.png)

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- Goose (for database migrations)

## Configuration

The service can be configured using environment variables:

```env
# Service Configuration
HOST=:50051

# Database Configuration
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_NAME=auth
POSTGRES_SSL_MODE=disable

# JWT Configuration
JWT_TTL=24
JWT_SECRET=your-secret-key

# Google OAuth2 Configuration
GOOGLE_CLIENT_ID=your-client-id
GOOGLE_CLIENT_SECRET=your-client-secret
GOOGLE_REDIRECT_URL=http://localhost:3000/auth/google/callback

# GitHub OAuth2 Configuration
GITHUB_CLIENT_ID=your-client-id
GITHUB_CLIENT_SECRET=your-client-secret
GITHUB_REDIRECT_URL=http://localhost:3000/auth/github/callback

# Logging Configuration
LOG_DEV=true
LOG_LEVEL=debug
LOG_FILE_OUTPUT=false
LOG_FILE_PATH=./logs/auth.log
LOG_STACKTRACE=true
LOG_CALLER=true
```

## Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up the database:
   ```bash
   # Run migrations
   goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/auth?sslmode=disable" up
   ```
4. Configure environment variables (see Configuration section)
5. Run the service:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Authentication

- `POST /auth/register` - Register a new user
- `POST /auth/login` - Login with email/password
- `GET /auth/google` - Initiate Google OAuth2 flow
- `GET /auth/github` - Initiate GitHub OAuth2 flow
- `POST /auth/refresh` - Refresh JWT token
- `POST /auth/logout` - Logout user

### Company Management

- `POST /companies` - Create a new company
- `GET /companies` - List companies
- `GET /companies/{id}` - Get company details
- `POST /companies/{id}/participants` - Add participant to company
- `DELETE /companies/{id}/participants/{userId}` - Remove participant from company

### User Management

- `GET /users/{id}` - Get user details
- `PUT /users/{id}` - Update user profile
- `POST /users/password/reset` - Request password reset
- `POST /users/password/change` - Change password

## Development

### Running Tests

```bash
go test ./...
```

### Database Migrations

```bash
# Create a new migration
goose -dir ./migrations create migration_name sql

# Apply migrations
goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/auth?sslmode=disable" up

# Rollback migrations
goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/auth?sslmode=disable" down
```
