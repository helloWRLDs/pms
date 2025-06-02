# Notifier Service

The Notifier Service is a crucial component of the PMS (Project Management System) responsible for handling all notification-related functionality. It manages email notifications, message queue processing, and ensures reliable delivery of notifications to users.

## Features

- Email Notifications
  - Task assignment notifications
  - Project updates
  - Sprint status changes
  - Document sharing notifications
  - System alerts
- Message Queue Integration
  - Reliable message processing
  - Message persistence
  - Retry mechanisms
  - Dead letter queue handling
- Template Management
  - HTML email templates
  - Dynamic content rendering
  - Multi-language support
- Notification Tracking
  - Delivery status tracking
  - Read receipts
  - Notification preferences

## Architecture

The service follows a modular architecture with the following components:

- **Handlers**: gRPC handlers for service endpoints
- **Service**: Core notification service implementation
- **Modules**:
  - Email module for email notifications
  - Message queue module for reliable message processing
- **Config**: Service configuration management

## Prerequisites

- Go 1.21 or higher
- RabbitMQ 3.9 or higher
- SMTP server (Gmail SMTP supported)
- Goose (for database migrations)

## Configuration

The service can be configured using environment variables:

```env
# Service Configuration
HOST=:50054

# Gmail SMTP Configuration
GMAIL_HOST=smtp.gmail.com
GMAIL_USERNAME=your-email@gmail.com
GMAIL_PASSWORD=your-app-specific-password
GMAIL_PORT=587

# RabbitMQ Configuration
MQ_DSN=amqp://guest:guest@localhost:5672/
MQ_EXCHANGE=notifications
MQ_DISABLE_LOG=false

# Logging Configuration
LOG_DEV=true
LOG_LEVEL=debug
LOG_FILE_OUTPUT=false
LOG_FILE_PATH=./logs/notifier.log
LOG_STACKTRACE=true
LOG_CALLER=true
```

## Setup

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Set up RabbitMQ:
   ```bash
   # Create exchange
   rabbitmqadmin declare exchange name=notifications type=direct
   ```
4. Configure environment variables (see Configuration section)
5. Run the service:
   ```bash
   go run cmd/main.go
   ```

## API Endpoints

### Notification Management

- `POST /notifications` - Send a new notification
- `GET /notifications` - List notifications
- `GET /notifications/{id}` - Get notification details
- `PUT /notifications/{id}/status` - Update notification status
- `DELETE /notifications/{id}` - Delete notification

### Email Templates

- `POST /templates` - Create a new email template
- `GET /templates` - List email templates
- `GET /templates/{id}` - Get template details
- `PUT /templates/{id}` - Update template
- `DELETE /templates/{id}` - Delete template

### Notification Preferences

- `GET /preferences` - Get user notification preferences
- `PUT /preferences` - Update notification preferences
- `POST /preferences/email` - Update email notification settings
- `POST /preferences/in-app` - Update in-app notification settings

## Message Queue Structure

### Exchanges

- `notifications` (direct) - Main exchange for all notifications

### Queues

- `email-notifications` - Queue for email notifications
- `in-app-notifications` - Queue for in-app notifications
- `notification-dlq` - Dead letter queue for failed notifications

### Routing Keys

- `email.*` - Email notification routing
- `in-app.*` - In-app notification routing
- `system.*` - System notification routing

## Email Templates

The service supports HTML email templates with the following features:

- Dynamic content insertion
- Conditional rendering
- Multi-language support
- Responsive design
- Custom styling

Example template structure:

```html
<!DOCTYPE html>
<html>
<head>
    <title>{{.Title}}</title>
</head>
<body>
    <h1>{{.Header}}</h1>
    <div class="content">
        {{.Content}}
    </div>
    <footer>
        {{.Footer}}
    </div>
</body>
</html>
```

## Development

### Running Tests

```bash
go test ./...
```

### Local Development

1. Start RabbitMQ:

   ```bash
   docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management
   ```

2. Start the service in development mode:
   ```bash
   go run cmd/main.go
   ```

## Security Considerations

- SMTP credentials are stored securely
- Message queue connections use TLS
- Email content is sanitized
- Rate limiting is implemented
- Notification preferences are enforced
- Sensitive data is encrypted

## Monitoring

The service provides the following monitoring capabilities:

- Message queue metrics
- Email delivery statistics
- Template rendering performance
- Error rates and types
- Queue lengths and processing times
