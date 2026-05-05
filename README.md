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
| **CLI Tool** | **Gons CLI** for scaffolding and project management |
| **Hot Reload** | Supported via [Air](https://github.com/air-verse/air) |

---

## [>] Installation & Quick Start

### 1. Install GONS CLI

To make development easier, install the Gons CLI tool globally:

```bash
go install github.com/rendrabagasdev/gons_cli@latest
```

*Make sure your `$GOPATH/bin` is in your system `PATH`.*

### 2. Create a New Project

```bash
gons new my-project
cd my-project
```

### 3. Setup Environment

```bash
cp .env .env.local  # Edit .env to match your database settings
npm install
go mod tidy
```

### 4. Run the Application

| Command | Description |
|---|---|
| `npm run dev` | Run Vite development server |
| `npm run air` | Run Go server with hot reload |
| `npm run migrate` | Run database migrations |
| `npm run seed` | Run database seeders |

---

## [!] CLI Usage

The Gons CLI provides several commands to speed up your workflow:

```bash
# Generate a new Controller
gons make:controller UserController

# Generate a new Model
gons make:model User

# Generate a new Service
gons make:service UserService

# Generate a new Request (Validation)
gons make:request UserRequest

# Generate a new Middleware
gons make:middleware AuthMiddleware
```

---

## [/] Directory Structure

```
gons/
├── app/
│   ├── http/
│   │   ├── controllers/    # HTTP controllers
│   │   ├── middlewares/    # HTTP middlewares
│   │   ├── requests/       # Request validation structs
│   │   └── services/       # Business logic
│   ├── models/             # GORM models
│   └── utils/              # Core utilities (Cache, Queue, Storage)
├── config/                 # Configurations
├── database/               # Migrations & Seeders
├── public/                 # Static files (Vite output)
├── resources/
│   └── views/              # HTML templates
├── routes/                 # Route definitions
├── .air.toml               # Hot reload config
├── main.go                 # Entry point
└── vite.config.js          # Vite configuration
```

---

## [=] Templating (HTML + Vite)

HTML templates are stored in `resources/views/`. Use the `vite` helper in templates to include Vite-managed assets:

```html
<!-- resources/views/layouts/main.html -->
<head>
    {{ vite "resources/js/app.js" }}
</head>
```

The `vite` helper automatically detects whether to use the dev server or built assets via the `public/hot` detection mechanism.

---

## [i] Tech Stack

- [Go](https://go.dev/) — primary programming language
- [Fiber v3](https://gofiber.io/) — web framework
- [GORM](https://gorm.io/) — ORM
- [Vite](https://vitejs.dev/) — frontend build tool
- [Air](https://github.com/air-verse/air) — hot reload
- [Gons CLI](https://github.com/rendrabagasdev/gons_cli) — scaffolding tool

---

## [c] License

This project is licensed under the [MIT License](LICENSE).
