# Analytics Service

The Analytics Service is a Go-based microservice responsible for generating and managing analytics reports for the Project Management System (PMS).

## Overview

This service provides functionality for:

- Generating PDF reports using wkhtmltopdf
- Processing and analyzing project data
- Managing analytics-related operations

## Prerequisites

- Go 1.23.2 or higher
- wkhtmltopdf (required for PDF generation)

## Project Structure

```
analytics/
├── cmd/            # Application entry points
├── internal/       # Private application code
│   ├── clients/    # External service clients
│   ├── config/     # Configuration management
│   ├── data/       # Data models and repositories
│   ├── handlers/   # HTTP handlers
│   ├── logic/      # Business logic
│   ├── modules/    # Feature modules
│   └── utils/      # Utility functions
├── migrations/     # Database migrations
└── bin/           # Binary output directory
```

## Dependencies

- github.com/adrg/go-wkhtmltopdf v0.3.0
- github.com/SebastiaanKlippert/go-wkhtmltopdf v1.9.3
- github.com/stretchr/testify v1.10.0 (for testing)

## Getting Started

1. Install dependencies:

   ```bash
   go mod download
   ```

2. Install wkhtmltopdf:

   - Windows: Download from https://wkhtmltopdf.org/downloads.html
   - Linux: `sudo apt-get install wkhtmltopdf`
   - macOS: `brew install wkhtmltopdf`

3. Run the service:
   ```bash
   go run cmd/main.go
   ```

## Development

To run tests:

```bash
go test ./...
```
