# Project Service

The Project Service is a core component of the PMS (Project Management System) responsible for managing projects, tasks, sprints, and related resources. It provides comprehensive project management functionality including task tracking, sprint planning, and project analytics.

## Features

- Project Management
  - Project creation and configuration
  - Project status tracking
  - Project code and prefix management
  - Project analytics and statistics
- Task Management
  - Task creation and assignment
  - Task status tracking
  - Task priority management
  - Task type categorization
- Sprint Management
  - Sprint planning and creation
  - Sprint status tracking
- Document Management
  - Document creation and versioning
  - Document sharing and collaboration
  - Document status tracking

## Architecture

The service follows a clean architecture pattern with the following components:

- **Handlers**: gRPC handlers for service endpoints
- **Logic**: Business logic implementation
- **Data**: Data access layer with PostgreSQL
- **Config**: Service configuration management

## Prerequisites

- Go 1.21 or higher
- PostgreSQL 14 or higher
- Goose (for database migrations)

## Configuration

The service can be configured using environment variables:

```env
# Service Configuration
HOST=:50052

# Database Configuration
POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_USER=postgres
POSTGRES_PASSWORD=postgres
POSTGRES_NAME=project
POSTGRES_SSL_MODE=disable

# Logging Configuration
LOG_DEV=true
LOG_LEVEL=debug
LOG_FILE_OUTPUT=false
LOG_FILE_PATH=./logs/project.log
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
   goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/project?sslmode=disable" up
   ```
4. Configure environment variables (see Configuration section)
5. Run the service:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Project Management

- `POST /projects` - Create a new project
- `GET /projects` - List projects
- `GET /projects/{id}` - Get project details
- `PUT /projects/{id}` - Update project
- `DELETE /projects/{id}` - Delete project

### Task Management

- `POST /tasks` - Create a new task
- `GET /tasks` - List tasks
- `GET /tasks/{id}` - Get task details
- `PUT /tasks/{id}` - Update task
- `DELETE /tasks/{id}` - Delete task
- `POST /tasks/{id}/assign` - Assign task to user
- `POST /tasks/{id}/unassign` - Unassign task from user

### Sprint Management

- `POST /sprints` - Create a new sprint
- `GET /sprints` - List sprints
- `GET /sprints/{id}` - Get sprint details
- `PUT /sprints/{id}` - Update sprint
- `DELETE /sprints/{id}` - Delete sprint
- `POST /sprints/{id}/start` - Start sprint
- `POST /sprints/{id}/complete` - Complete sprint

### Document Management

- `POST /documents` - Create a new document
- `GET /documents` - List documents
- `GET /documents/{id}` - Get document details
- `PUT /documents/{id}` - Update document
- `DELETE /documents/{id}` - Delete document
- `POST /documents/{id}/share` - Share document
- `POST /documents/{id}/version` - Create new version

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
goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/project?sslmode=disable" up

# Rollback migrations
goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/project?sslmode=disable" down
```

## Data Models

### Project

- ID (UUID)
- Title
- Description
- CompanyID
- Status
- CodeName
- CodePrefix
- CreatedAt
- UpdatedAt

### Task

- ID (UUID)
- Title
- Description
- Status
- Priority
- ProjectID
- AssigneeID
- Type
- DueDate
- CreatedAt
- UpdatedAt

### Sprint

- ID (UUID)
- Title
- Description
- ProjectID
- StartDate
- EndDate
- Status
- CreatedAt
- UpdatedAt

### Document

- ID (UUID)
- Title
- Content
- ProjectID
- Status
- Version
- CreatedBy
- CreatedAt
- UpdatedAt
