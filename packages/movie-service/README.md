# Movie Service

## Overview

The Movie Service is a microservice responsible for managing movie data from the IMDb dataset.
It provides REST API endpoints for searching, retrieving, and filtering movie information.

## Technology Stack

- **Language**: Go
- **Framework**: Gin Web Framework
- **Database**: MongoDB
- **Docker**: Containerization

## Features

- Movie data import from IMDb TSV datasets
- CRUD operations for movie entities
- Search functionality with various filters
- Pagination support
- Health check endpoint

## Directory Structure

```
movie-service/
├── cmd/
│   └── api/
│       └── main.go               # Application entry point
│
├── internal/
│   ├── handlers/                 # API endpoint handlers
│   │   ├── movies.go             # Movie-related endpoints
│   │   └── health.go             # Health check endpoint
│   │
│   ├── models/                   # Data models
│   │   ├── movie.go              # Movie entity model
│   │   └── errors.go             # Custom error definitions
│   │
│   ├── repositories/             # Data access layer
│   │   └── movie_repository.go   # Movie data repository
│   │
│   └── services/                 # Business logic
│       └── movie_service.go      # Movie-related business operations
│
├── pkg/
│   ├── db/                       # Database connections
│   │   └── mongodb.go            # MongoDB client setup
│   │
│   └── parser/                   # TSV parsing utilities
│       ├── tsv_parser.go         # Parser implementation
│       └── parser_test.go        # Parser tests
│
├── configs/                      # Configuration files
│   └── config.go                 # Configuration loading
│
├── go.mod                        # Go module definition
└── go.sum                        # Dependency checksums
```

## Setup and Running

### Prerequisites

- Go 1.19+
- MongoDB
- IMDb dataset files (see main README for download instructions)

### Environment Variables

```
MONGO_URI=mongodb://localhost:27017
MONGO_DB=imdb
PORT=8080
LOG_LEVEL=info
```

### How to Run

#### 1. Install dependencies

```bash
go mod download
```

#### 2. Run locally

```bash
go run cmd/api/main.go
```

#### 3. Build

```bash
go build -o movie-service cmd/api/main.go
```

#### 4. Run with Docker

```bash
# Build Docker image
docker build -t movie-service -f ../../docker/services/movie-service.dockerfile .

# Run container
docker run -p 8080:8080 movie-service
```

## API Endpoints

### Movies

#### GET /api/v1/movies

Get a list of movies with pagination

Query Parameters:

- `page`: Page number (default: 1)
- `limit`: Items per page (default: 10)
- `sort`: Sort field (default: "title")
- `order`: Sort order, "asc" or "desc" (default: "asc")
- `title`: Filter by title (partial match)
- `year`: Filter by release year
- `genres`: Filter by genres (comma-separated)

#### GET /api/v1/movies/:id

Get a movie by ID

#### GET /api/v1/movies/search

Search movies with advanced filtering

#### POST /api/v1/movies

Create a new movie (admin only)

#### PUT /api/v1/movies/:id

Update a movie (admin only)

#### DELETE /api/v1/movies/:id

Delete a movie (admin only)

### Health

#### GET /health

Health check endpoint

## Data Import

The service provides functionality to import IMDb dataset files:

```bash
go run cmd/api/main.go --import-data
```

## Testing

Run tests with:

```bash
go test ./...
```

## Development Notes

- The service follows clean architecture principles
- Handlers are responsible for HTTP request/response handling
- Services contain business logic
- Repositories handle data access
- Models define data structures
