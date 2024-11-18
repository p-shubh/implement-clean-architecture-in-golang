# Clean Architecture in Golang

## Overview

This project demonstrates how to implement **Clean Architecture** in a Golang application. The structure separates the application into distinct layers, adhering to the **SOLID principles** for better maintainability, scalability, and testing.

The application includes a use case for tracking cockroach detections, with PostgreSQL as the database and Echo as the HTTP framework. The repository handles database interactions, while the use case implements core business logic.

## Features

- **Clean Architecture**: Separation of concerns into handler, domain (use case and repository), and infrastructure layers.
- **SOLID Principles**: Applied to ensure single responsibility, dependency inversion, and flexibility for testing.
- **PostgreSQL Integration**: With GORM for database ORM.
- **Dockerized PostgreSQL**: Simplified local development setup.
- **HTTP Server**: Built using Echo framework.
- **Config Management**: Simplified with Viper.

## Tech Stack

- Golang
- Echo
- GORM
- PostgreSQL
- Docker
- Viper

## Project Structure

```
📂 config/
├─ 📄 config.go
📂 server/
├─ 📄 echoServer.go
📂 database/
├─ 📄 postgres.go
📂 cockroach/
├─ 📂 entities/
│  ├─ 📄 cockroachEntity.go
├─ 📂 migrations/
│  ├─ 📄 cockroachMigration.go
├─ 📂 repositories/
│  ├─ 📄 cockroachRepository.go
├─ 📂 usecases/
│  ├─ 📄 cockroachUsecase.go
├─ 📂 handlers/
│  ├─ 📄 cockroachHandler.go
📄 main.go
📄 config.yaml
```

## Setup Instructions

### Prerequisites

- Install [Docker](https://www.docker.com/)
- Install [Golang](https://golang.org/)

### Steps

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/your-repo.git
   cd your-repo
   ```
2. **Start PostgreSQL in Docker**:

   ```bash
   docker pull postgres:alpine
   docker run --name cockroachdb -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres:alpine
   ```
3. **Set Up Database**:

   ```bash
   docker exec -it cockroachdb bash
   psql -U postgres -c "CREATE DATABASE cockroachdb;"
   ```
4. **Install Dependencies**:

   ```bash
   go mod tidy
   ```
5. **Run Migrations**:

   ```bash
   go run ./cockroach/migrations/cockroachMigration.go
   ```
6. **Run the Server**:

   ```bash
   go run main.go
   ```
7. **Test the API**:
   Use a tool like Postman or cURL:

   ```bash
   curl --location 'http://localhost:8080/v1/cockroach' \
   --header 'Content-Type: application/json' \
   --data '{
       "amount": 3
   }'
   ```

## Example Request and Response

- **Endpoint**: `POST /v1/cockroach`
- **Request**:
  ```json
  {
    "amount": 3
  }
  ```
- **Response**:
  ```json
  {
    "message": "Success 🪳🪳🪳"
  }
  ```

## Key Concepts

- **Handler Layer**: Validates and processes incoming requests.
- **Domain Layer**: Contains core business logic (use cases).
- **Repository Layer**: Manages interactions with external services like databases.

## Future Enhancements

- Add unit tests for each layer.
- Extend use cases for additional functionality.
- Implement more robust validation and error handling.

---

Let me know if you'd like to customize this further!
