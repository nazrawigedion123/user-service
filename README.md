# User Service

## Overview

This is a microservice responsible for managing user data within the VTS (Vessel Tracking System) ecosystem. It provides a RESTful API for creating, reading, updating, and deleting user records. The service is built with Go and follows hexagonal architecture principles to ensure separation of concerns, maintainability, and testability.

## Current Stage: Development

The service is currently in the **development stage**. The core functionalities for user management (CRUD) are implemented and functional. However, it is not yet production-ready.

**Key implemented features:**
*   Full CRUD operations for users.
*   Password hashing using `bcrypt`.
*   Soft-delete mechanism.
*   REST API with JSON request/response.
*   Structured logging with `zap`.
*   PostgreSQL database integration using `sqlc` for type-safe queries.

**Areas for future development:**
*   Authentication and authorization (e.g., JWT).
*   Comprehensive input validation.
*   Pagination for user lists.
*   Robust error handling and standardized error responses.
*   Unit and integration tests.
*   Containerization (Docker).
*   CI/CD pipeline setup.

## Architecture

The service is structured using a layered, clean architecture approach:

*   **Handlers (`internal/adapters/handlers`):** The outermost layer, responsible for handling HTTP requests and responses. It uses the Gin web framework. It decodes requests, calls the appropriate application service method, and formats the response.
*   **Application Service (`internal/application`):** Contains the core business logic. It orchestrates the flow of data and operations, independent of the delivery mechanism (HTTP) or the database technology.
*   **Domain (`internal/domain`):** Defines the core entities (e.g., `User`) and repository interfaces (`UserRepository`). This is the heart of the application.
*   **Repository (`internal/adapters/repository`):** Implements the `UserRepository` interface, providing a concrete data access layer for PostgreSQL. It uses `sqlc` to generate Go code from SQL queries.

## API Endpoints

| Method | Endpoint                       | Description                                |
| :----- | :----------------------------- | :----------------------------------------- |
| `POST` | `/users`                       | Creates a new user.                        |
| `GET`    | `/users`                       | Retrieves a list of all users.             |
| `GET`    | `/users/:id`                   | Retrieves a single user by their UUID.     |
| `GET`    | `/users/username/:username`    | Retrieves a single user by their username. |
| `GET`    | `/users/email/:email`          | Retrieves a single user by their email.    |
| `PUT`    | `/users`                       | Updates an existing user's information.    |
| `DELETE` | `/users/:id`                   | Soft-deletes a user by their UUID.         |

## Technologies Used

*   **Language:** Go
*   **Web Framework:** Gin
*   **Database:** PostgreSQL
*   **ORM/Query Builder:** sqlc
*   **Logging:** Zap
*   **UUID Generation:** google/uuid

## Setup and Running

### Prerequisites

*   Go (version 1.18 or higher)
*   PostgreSQL
*   `sqlc` CLI (for regenerating database code)

### Environment Variables

The service requires the following environment variables to be set:

*   `DB_SOURCE`: The connection string for the PostgreSQL database (e.g., `postgresql://root:secret@localhost:5432/user_db?sslmode=disable`).
*   `SALT`: A secret string used for hashing passwords.

### Installation

1.  **Clone the repository:**
    ```sh
    git clone <repository-url>
    cd user-service
    ```

2.  **Install dependencies:**
    ```sh
    go mod tidy
    ```

3.  **Set up the database:**
    Connect to your PostgreSQL instance and run the schema definition from `sql/schema/schema.sql` to create the `users` table and its indexes.

4.  **(Optional) Regenerate SQLC code:**
    If you make changes to the queries in `sql/query/queries.sql`, regenerate the Go code:
    ```sh
    sqlc generate
    ```

### Running the Service

1.  **Set environment variables:**
    ```sh
    export DB_SOURCE="postgresql://user:password@localhost:5432/users?sslmode=disable"
    export SALT="your-secret-salt-here"
    ```

2.  **Run the application:**
    ```sh
    go run ./cmd/main.go
    ```
    The service will start, typically on port `8080`.
