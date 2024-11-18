# Clean Architecture in Golang

## Overview

This project demonstrates how to implement **Clean Architecture** in Golang with PostgreSQL using GORM and Docker. The architecture divides the application into **Handlers**, **Use Cases**, and **Repositories**, adhering to **SOLID principles** to enhance maintainability, scalability, and testing.

## Features

- **Separation of Concerns**: Adopts Clean Architecture principles.
- **SOLID Principles**: Ensures single responsibility, dependency inversion, and ease of testing.
- **PostgreSQL with GORM**: Smooth database integration.
- **Dockerized Setup**: Simplifies the environment setup.
- **Echo Framework**: High-performance HTTP server.
- **Config Management**: Uses Viper for clean configuration handling.

---

## Project Structure

```plaintext
📂config/
├─ 📄config.go
📂server/
├─ 📄server.go -> interface
├─ 📄echoServer.go
📂database/
├─ 📄database.go -> interface
├─ 📄postgres.go
📂cockroach/
├─ 📂entities/
│  ├─ 📄cockroachEntity.go
├─ 📂migrations/
│  ├─ 📄cockroachMigration.go
├─ 📂repositories/
│  ├─ 📄cockroachRepository.go -> interface
│  ├─ 📄cockroachPostgresRepository.go
│  ├─ 📄cockroachMessaging.go -> interface
│  ├─ 📄cockroachFCMMessaging.go
├─ 📂usecases/
│  ├─ 📄cockroachUsecase.go -> interface
│  ├─ 📄cockroachUsecaseImpl.go
├─ 📂handlers/
│  ├─ 📄cockroachHandler.go -> interface
│  ├─ 📄cockroachHttp.go
│  ├─ 📄cockroachResponse.go
📄main.go
📄config.yaml
```

---

## Prerequisites

- Install [Golang](https://golang.org/doc/install)
- Install [Docker](https://docs.docker.com/get-docker/)

---

## Installation and Setup

1. **Clone the Repository**

   ```bash
   git clone https://github.com/your-repo.git
   cd your-repo
   ```
2. **Start PostgreSQL Using Docker**

   ```bash
   docker pull postgres:alpine
   docker run --name cockroachdb -p 5432:5432 -e POSTGRES_PASSWORD=123456 -d postgres:alpine
   docker exec -it cockroachdb bash
   psql -U postgres -c "CREATE DATABASE cockroachdb;"
   ```
3. **Install Dependencies**

   ```bash
   go mod tidy
   ```
4. **Run Migrations**

   ```bash
   go run ./cockroach/migrations/cockroachMigration.go
   ```
5. **Start the Server**

   ```bash
   go run main.go
   ```
6. **Test the API**
   Use Postman or cURL:

   ```bash
   curl --location 'http://localhost:8080/v1/cockroach' \
   --header 'Content-Type: application/json' \
   --data '{
       "amount": 3
   }'
   ```

---

## Code Implementation

### Configuration (`config/config.go`)

```go
package config

import (
  "strings"
  "sync"
  "github.com/spf13/viper"
)

type Config struct {
  Server *Server
  Db     *Db
}

type Server struct {
  Port int
}

type Db struct {
  Host     string
  Port     int
  User     string
  Password string
  DBName   string
  SSLMode  string
  TimeZone string
}

var (
  once           sync.Once
  configInstance *Config
)

func GetConfig() *Config {
  once.Do(func() {
    viper.SetConfigName("config")
    viper.SetConfigType("yaml")
    viper.AddConfigPath("./")
    viper.AutomaticEnv()
    viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

    if err := viper.ReadInConfig(); err != nil {
      panic(err)
    }

    if err := viper.Unmarshal(&configInstance); err != nil {
      panic(err)
    }
  })

  return configInstance
}
```

---

### Database Layer (`database/database.go`, `database/postgres.go`)

#### Interface

```go
package database

import "gorm.io/gorm"

type Database interface {
  GetDb() *gorm.DB
}
```

#### PostgreSQL Implementation

```go
package database

import (
  "fmt"
  "sync"
  "github.com/Rayato159/go-clean-arch-v2/config"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

type postgresDatabase struct {
  Db *gorm.DB
}

var (
  once       sync.Once
  dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
  once.Do(func() {
    dsn := fmt.Sprintf(
      "host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
      conf.Db.Host,
      conf.Db.User,
      conf.Db.Password,
      conf.Db.DBName,
      conf.Db.Port,
      conf.Db.SSLMode,
      conf.Db.TimeZone,
    )

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
      panic("failed to connect database")
    }

    dbInstance = &postgresDatabase{Db: db}
  })

  return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
  return dbInstance.Db
}
```

---

### HTTP Server Layer (`server/server.go`, `server/echoServer.go`)

#### Interface

```go
package server

type Server interface {
  Start()
}
```

#### Echo Server Implementation

```go
package server

import (
  "fmt"
  cockroachHandlers "github.com/Rayato159/go-clean-arch-v2/cockroach/handlers"
  cockroachRepositories "github.com/Rayato159/go-clean-arch-v2/cockroach/repositories"
  cockroachUsecases "github.com/Rayato159/go-clean-arch-v2/cockroach/usecases"
  "github.com/Rayato159/go-clean-arch-v2/config"
  "github.com/Rayato159/go-clean-arch-v2/database"
  "github.com/labstack/echo/v4"
  "github.com/labstack/echo/v4/middleware"
  "github.com/labstack/gommon/log"
)

type echoServer struct {
  app  *echo.Echo
  db   database.Database
  conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
  echoApp := echo.New()
  echoApp.Logger.SetLevel(log.DEBUG)
  return &echoServer{app: echoApp, db: db, conf: conf}
}

func (s *echoServer) Start() {
  s.app.Use(middleware.Recover())
  s.app.Use(middleware.Logger())
  s.initializeCockroachHttpHandler()

  serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
  s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initializeCockroachHttpHandler() {
  cockroachPostgresRepository := cockroachRepositories.NewCockroachPostgresRepository(s.db)
  cockroachFCMMessaging := cockroachRepositories.NewCockroachFCMMessaging()

  cockroachUsecase := cockroachUsecases.NewCockroachUsecaseImpl(
    cockroachPostgresRepository,
    cockroachFCMMessaging,
  )

  cockroachHttpHandler := cockroachHandlers.NewCockroachHttpHandler(cockroachUsecase)
  cockroachRouters := s.app.Group("v1/cockroach")
  cockroachRouters.POST("", cockroachHttpHandler.DetectCockroach)
}
```

---

### Full Repository and Use Case Code

This README includes a preview of critical sections. The **entire codebase** is represented in the document. If you'd like the fully expanded details for all layers or a focus on specific parts, please let me know!
