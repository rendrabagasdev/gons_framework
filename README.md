<div align="center">

```
   ______                 
  / ____/___  ____  _____ 
 / / __/ __ \/ __ \/ ___/ 
/ /_/ / /_/ / / / (__  )  
\____/\____/_/ /_/____/   
```

# Gons Framework

**Go web framework starter kit — powered by Fiber, GORM, and a clean layered architecture.**

[![Go Version](https://img.shields.io/badge/Go-1.21%2B-00ADD8?logo=go)](https://go.dev/)
[![Fiber](https://img.shields.io/badge/Fiber-v3-00ACD7?logo=go)](https://gofiber.io/)
[![GORM](https://img.shields.io/badge/GORM-v2-lightgrey)](https://gorm.io/)
[![License](https://img.shields.io/badge/License-MIT-green)](LICENSE)

</div>

---

## [+] Key Features

| Feature | Description |
|---|---|
| **Web Framework** | [Fiber v3](https://gofiber.io/) — fast HTTP framework built on Fasthttp |
| **ORM** | [GORM](https://gorm.io/) — supports MySQL, PostgreSQL, and SQLite |
| **Dependency Injection** | [Golobby Container v3](https://github.com/golobby/container) |
| **Template Engine** | HTML templating integrated with [Vite](https://vitejs.dev/) |
| **Cache** | Memory driver (ready to extend with Redis) |
| **Queue** | Goroutine/Sync driver (ready to extend with Redis) |
| **Storage** | Local & S3-compatible drivers (AWS S3, Supabase, MinIO) |
| **Middleware** | Auth Guard, CORS, Logger, Static Files |
| **Migration & Seeder** | GORM auto-migrate + seeder via CLI |
| **Hot Reload** | Supported via [Air](https://github.com/air-verse/air) |

---

## [/] Directory Structure

```
gons/
├── app/
│   ├── contracts/          # Interfaces / contracts (Cache, Queue, Storage)
│   ├── http/
│   │   ├── controllers/    # HTTP controllers
│   │   ├── middlewares/    # HTTP middlewares (AuthGuard, etc.)
│   │   ├── requests/       # Request validation structs
│   │   └── services/       # Business logic
│   ├── models/             # GORM models & model registry
│   └── utils/
│       ├── cache/          # Cache driver implementations
│       ├── queue/          # Queue driver implementations
│       ├── storage/        # Storage driver implementations (Local & S3)
│       └── validator/      # Request validation helpers
├── bootstrap/
│   ├── app.go              # Application bootstrap
│   └── migrate.go          # Migration command
├── config/
│   ├── cache.go            # Cache configuration
│   ├── database.go         # Database configuration
│   ├── env.go              # Environment variable helper
│   ├── queue.go            # Queue configuration
│   ├── registry.go         # Config provider registry
│   ├── server.go           # Fiber server configuration
│   ├── service.go          # DI service configuration
│   ├── storage.go          # Storage configuration
│   └── vite.go             # Vite integration helper
├── database/
│   └── seeders/            # Database seeders
├── public/                 # Static files (CSS, JS, images)
├── resources/
│   └── views/              # HTML templates
├── routes/
│   ├── registry.go         # Route group registration
│   └── web.go              # Web route definitions
├── .air.toml               # Air hot reload configuration
├── .env                    # Environment variables
├── go.mod
├── main.go                 # Application entry point
├── package.json            # Frontend dependencies (Vite)
└── vite.config.js          # Vite configuration
```

---

## [>] Quick Start

### Prerequisites

- [Go](https://go.dev/dl/) 1.21+
- [Node.js](https://nodejs.org/) 18+ (for frontend assets)
- A supported database: MySQL, PostgreSQL, or SQLite

### Installation

```bash
# Clone the repository
git clone https://github.com/rendrabagasdev/gons.git
cd gons

# Install Go dependencies
go mod tidy

# Install Node.js dependencies (frontend / Vite)
npm install

# Copy the environment configuration file
cp .env .env.local  # edit .env to match your environment
```

### `.env` Configuration

```env
#### DATABASE ENV ####
DB_DRIVER=sqlite          # mysql | postgres | sqlite
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=
DB_NAME=gons_db

#### SERVER ENV ####
APP_NAME=Gons Framework
APP_PORT=8080
APP_HOST=localhost
APP_ENV=development
APP_URL=http://localhost:8080
APP_TIMEZONE=Asia/Jakarta

#### STORAGE ####
STORAGE_DISK=local        # local | s3

#### CACHE ####
CACHE_DRIVER=memory       # memory | redis

#### QUEUE ####
QUEUE_DRIVER=sync         # sync | goroutine

#### CORS ####
CORS_ALLOW_ORIGINS=*
CORS_ALLOW_HEADERS=Origin,Content-Type,Accept,Authorization
```

### Run Migrations

```bash
go run main.go migrate
```

### Run Seeders

```bash
go run main.go --seed
```

### Run the Server

```bash
# Production
go run main.go

# Development with hot reload (requires Air)
go install github.com/air-verse/air@latest
air
```

The server will be available at `http://localhost:8080`.

---

## [~] Routing

Route definitions live in the `routes/` directory. Add new routes in `routes/web.go`:

```go
func WebRoutes(app *fiber.App) {
    ctrl := controllers.NewController()
    userCtrl := controllers.NewUserController()

    // Public routes
    app.Get("/", ctrl.Welcome)
    app.Get("/users", userCtrl.GetAll)

    // Protected routes (require Authorization header)
    protected := app.Group("/protected")
    protected.Use(middlewares.AuthGuard())
    protected.Get("/users", userCtrl.GetAll)
}
```

---

## [#] Creating a Controller

```go
// app/http/controllers/PostController.go
package controllers

import (
    "github.com/gofiber/fiber/v3"
    "github.com/golobby/container/v3"
)

type PostController struct{}

func NewPostController() *PostController {
    ctrl := &PostController{}
    container.Fill(ctrl)
    return ctrl
}

func (c *PostController) Index(ctx fiber.Ctx) error {
    return ctx.JSON(fiber.Map{"message": "Hello from PostController"})
}
```

---

## [*] Database (GORM)

Register models in `app/models/registry.go`:

```go
var ModelRegistry = []interface{}{
    &User{},
    &Post{}, // add new models here
}
```

Run migrations:

```bash
go run main.go migrate
```

### Supported Drivers

| Driver | `DB_DRIVER` |
|---|---|
| MySQL | `mysql` |
| PostgreSQL | `postgres` |
| SQLite | `sqlite` |

---

## [-] Cache

Use the `contracts.Cache` interface via dependency injection:

```go
import "gons/app/contracts"

type MyService struct {
    Cache contracts.Cache `container:"type"`
}
```

Available drivers: `memory` (default), `redis` (coming soon).

---

## [-] Queue

Use the `contracts.Queue` interface via dependency injection:

```go
import "gons/app/contracts"

type MyService struct {
    Queue contracts.Queue `container:"type"`
}
```

Available drivers: `sync`, `goroutine`, `redis` (coming soon).

---

## [-] Storage

Use the `contracts.Storage` interface via dependency injection:

```go
import "gons/app/contracts"

type MyService struct {
    Storage contracts.Storage `container:"type"`
}
```

### S3 / Supabase / MinIO Configuration

```env
STORAGE_DISK=s3
S3_KEY=your-access-key
S3_SECRET=your-secret-key
S3_REGION=ap-southeast-1
S3_BUCKET=my-bucket
S3_ENDPOINT=https://your-endpoint.supabase.co/storage/v1/s3
S3_USE_PATH_STYLE=true
```

---

## [!] Middleware

### AuthGuard

The `AuthGuard` middleware checks for the presence of the `Authorization` header. Apply it to any route group:

```go
protected := app.Group("/api")
protected.Use(middlewares.AuthGuard())
```

### Adding Custom Middleware

Create a new file in `app/http/middlewares/` and implement `fiber.Handler`.

---

## [=] Templating (HTML + Vite)

HTML templates are stored in `resources/views/`. Use the `vite` helper in templates to include Vite-managed assets:

```html
<!-- resources/views/layouts/main.html -->
<head>
    {{ vite "resources/js/app.js" }}
</head>
```

Build frontend assets:

```bash
npm run build
```

---

## [i] Tech Stack

- [Go](https://go.dev/) — primary programming language
- [Fiber v3](https://gofiber.io/) — web framework
- [GORM](https://gorm.io/) — ORM
- [Golobby Container](https://github.com/golobby/container) — dependency injection
- [godotenv](https://github.com/joho/godotenv) — `.env` file management
- [go-playground/validator](https://github.com/go-playground/validator) — input validation
- [Vite](https://vitejs.dev/) — frontend build tool
- [Air](https://github.com/air-verse/air) — hot reload

---

## [c] License

This project is licensed under the [MIT License](LICENSE).
