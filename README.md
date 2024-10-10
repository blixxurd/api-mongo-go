# Experimental Go API with MongoDB

## Overview

This project is an exploration into building robust, scalable APIs using Go. This app is designed to explore and implement clean architecture principles, emphasizing separation of concerns and modularity.

## Key Features

- Clean Architecture implementation in Go using a controller/service/repository
- MongoDB integration
- Clean and composable RESTful API design
- Docker support for local MongoDB
- Experiments with Dependency injection and Inversion of Control patterns

## Project Structure

The project follows a clean opinionated architecture pattern, separating concerns into distinct layers:

```
api-mongo-go/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── controllers/             # HTTP request handlers
│   ├── models/                  # Database models
│   ├── core/                    # Core Application Providers
│   ├── repositories/            # Database access layer
│   ├── routes/                  # API route definitions
│   └── services/                # Business logic layer
├── pkg/
│   └── database/                # Database connection utilities
├── docker-compose.yml           # Docker setup for MongoDB
├── go.mod
└── README.md
```

### Layer Responsibilities

1. **Controllers**: Handle HTTP requests and responses. They depend on Services for business logic.
2. **Services**: Implement business logic. They depend on Repositories for data access.
3. **Repositories**: Manage data access to MongoDB. They are database-specific.
4. **Models**: Define data structures used throughout the application.
5. **Core**: Manage dependency injection providers, providing a clean way to pass database connections and other dependencies from the top level application layer

## Design Decisions

- **Dependency Injection**: Chosen to make it easier to provide top down interfaces to the core application, and aide in future testing. 
- **Repository Pattern**: Used to abstract database operations, allowing for easier testing and potential database changes.
- **Service Layer**: Introduced to encapsulate business logic, keeping controllers thin and focused on HTTP concerns.
- **MongoDB Integration**: Chosen to provide leniancy for flat data models, allowing me to focus on application structure. 
- **Docker Usage**: Leveraged to ensure consistent development environments.

## Reflections and Future Improvements

This project is intended to be an eternal work in progress, and will be updated as my opinions and experimentations see fit. Some areas I'm considering for future iterations:

1. Implementing comprehensive unit and integration tests
2. Implementing logging and monitoring solutions
3. Exploring gRPC alongside REST for API communication
4. Containerizing the Go application for a full Docker-compose setup

## Running the Project

1. Start MongoDB:
   ```
   docker-compose up -d
   ```

2. Run the application:
   ```
   go run cmd/main.go
   ```