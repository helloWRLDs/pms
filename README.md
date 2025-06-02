# Project Management System (PMS)

A modern, scalable, and feature-rich Project Management System built with microservices architecture. The system provides comprehensive project management capabilities, team collaboration, analytics, and real-time notifications.

## System Architecture

The PMS consists of the following microservices:

### Core Services

- **API Gateway** (`services/api-gateway`)

  - Central entry point for all client requests
  - Request routing and service orchestration
  - Authentication and authorization
  - Protocol translation (gRPC to HTTP)

- **Auth Service** (`services/auth`)

  - User authentication and authorization
  - Session management
  - OAuth2 integration
  - User profile management

- **Project Service** (`services/project`)

  - Project lifecycle management
  - Task management
  - Team collaboration
  - Document management

- **Analytics Service** (`services/analytics`)

  - Project metrics and KPIs
  - Performance analytics
  - Custom reports
  - Data visualization

- **Notifier Service** (`services/notifier`)
  - Real-time alerts
  - Message queue integration
  - Template management

### Frontend

- **Web Application** (`services/frontend`)
  - Modern React-based UI
  - Real-time updates
  - Responsive design
  - Rich user experience

## Technology Stack

### Backend

- Go 1.21+
- gRPC
- Redis
- RabbitMQ
- PostgreSQL
- Docker

### Frontend

- React 18
- TypeScript
- Tailwind CSS
- Vite
- TanStack Query
- Zustand

## Prerequisites

- Go 1.21 or higher
- Node.js 18 or higher
- Docker and Docker Compose
- PostgreSQL 15 or higher
- Redis 7.0 or higher
- RabbitMQ 3.9 or higher

## Quick Start

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/pms.git
   cd pms
   ```

2. Start the development environment:

   ```bash
   docker-compose up -d
   ```

3. Start individual services:

   ```bash
   # Start API Gateway
   make run service=api-gateway

   # Start Auth Service
   make run service=auth

   # Start Project Service
   make run service=project

   # Start Analytics Service
   make run service=analytics

   # Start Notifier Service
   make run service=notifier

   # Start Frontend
   cd services/frontend
   npm install
   npm run dev
   ```

## Development

### Project Structure

```
pms/
├── services/
│   ├── api-gateway/    # API Gateway Service
│   ├── auth/          # Authentication Service
│   ├── project/       # Project Management Service
│   ├── analytics/     # Analytics Service
│   ├── notifier/      # Notification Service
│   └── frontend/      # Web Application
├── pkg/               # Shared packages
└── docker-compose.yml # Development environment
```

### Environment Setup

1. Create `.env` files for each service (see individual service READMEs)
2. Configure database connections
3. Set up message queues
4. Configure service discovery

### Running Tests

```bash
# Run all tests
go test ./...

# Run frontend tests
cd services/frontend
npm test
```

## Deployment

### Docker Deployment

1. Build images:

   ```bash
   docker-compose build
   ```

2. Deploy services:
   ```bash
   docker-compose up -d
   ```
