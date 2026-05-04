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

## ✨ Fitur Utama

| Fitur | Keterangan |
|---|---|
| 🚀 **Web Framework** | [Fiber v3](https://gofiber.io/) — framework HTTP cepat berbasis Fasthttp |
| 🗄️ **ORM** | [GORM](https://gorm.io/) — mendukung MySQL, PostgreSQL, dan SQLite |
| 💉 **Dependency Injection** | [Golobby Container v3](https://github.com/golobby/container) |
| 🎨 **Template Engine** | HTML templating terintegrasi dengan [Vite](https://vitejs.dev/) |
| 💾 **Cache** | Driver Memory (siap diperluas ke Redis) |
| 📬 **Queue** | Driver Goroutine/Sync (siap diperluas ke Redis) |
| 🗃️ **Storage** | Driver Lokal & S3-compatible (AWS S3, Supabase, MinIO) |
| 🔐 **Middleware** | Auth Guard, CORS, Logger, Static Files |
| 🔄 **Migrasi & Seeder** | Auto-migrate GORM + seeder via CLI |
| 🔥 **Hot Reload** | Dukungan [Air](https://github.com/air-verse/air) |

---

## 📁 Struktur Direktori

```
gons/
├── app/
│   ├── contracts/          # Interface / kontrak (Cache, Queue, Storage)
│   ├── http/
│   │   ├── controllers/    # HTTP controllers
│   │   ├── middlewares/    # HTTP middlewares (AuthGuard, dll)
│   │   ├── requests/       # Struct validasi request
│   │   └── services/       # Business logic
│   ├── models/             # GORM models & model registry
│   └── utils/
│       ├── cache/          # Implementasi driver cache
│       ├── queue/          # Implementasi driver queue
│       ├── storage/        # Implementasi driver storage (Local & S3)
│       └── validator/      # Helper validasi request
├── bootstrap/
│   ├── app.go              # Bootstrap aplikasi
│   └── migrate.go          # Perintah migrasi
├── config/
│   ├── cache.go            # Konfigurasi cache
│   ├── database.go         # Konfigurasi database
│   ├── env.go              # Helper pembaca env
│   ├── queue.go            # Konfigurasi queue
│   ├── registry.go         # Config provider registry
│   ├── server.go           # Konfigurasi server Fiber
│   ├── service.go          # Konfigurasi service DI
│   ├── storage.go          # Konfigurasi storage
│   └── vite.go             # Integrasi Vite helper
├── database/
│   └── seeders/            # Database seeders
├── public/                 # File statis (CSS, JS, gambar)
├── resources/
│   └── views/              # Template HTML
├── routes/
│   ├── registry.go         # Pendaftaran semua route group
│   └── web.go              # Definisi rute web
├── .air.toml               # Konfigurasi hot reload Air
├── .env                    # Variabel lingkungan
├── go.mod
├── main.go                 # Entry point aplikasi
├── package.json            # Dependensi frontend (Vite)
└── vite.config.js          # Konfigurasi Vite
```

---

## 🚀 Mulai Cepat

### Prasyarat

- [Go](https://go.dev/dl/) 1.21+
- [Node.js](https://nodejs.org/) 18+ (untuk aset frontend)
- Database pilihan: MySQL, PostgreSQL, atau SQLite

### Instalasi

```bash
# Clone repositori
git clone https://github.com/rendrabagasdev/gons.git
cd gons

# Install dependensi Go
go mod tidy

# Install dependensi Node.js (frontend / Vite)
npm install

# Salin file konfigurasi lingkungan
cp .env .env.local  # sesuaikan isi .env sesuai kebutuhan
```

### Konfigurasi `.env`

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

### Menjalankan Migrasi

```bash
go run main.go migrate
```

### Menjalankan Seeder

```bash
go run main.go --seed
```

### Menjalankan Server

```bash
# Produksi
go run main.go

# Development dengan hot reload (membutuhkan Air)
go install github.com/air-verse/air@latest
air
```

Server akan berjalan di `http://localhost:8080`.

---

## 🗺️ Routing

Definisi rute terdapat di direktori `routes/`. Tambahkan rute baru di `routes/web.go`:

```go
func WebRoutes(app *fiber.App) {
    ctrl := controllers.NewController()
    userCtrl := controllers.NewUserController()

    // Rute publik
    app.Get("/", ctrl.Welcome)
    app.Get("/users", userCtrl.GetAll)

    // Rute terproteksi (membutuhkan header Authorization)
    protected := app.Group("/protected")
    protected.Use(middlewares.AuthGuard())
    protected.Get("/users", userCtrl.GetAll)
}
```

---

## 🏗️ Membuat Controller

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

## 🗄️ Database (GORM)

Daftarkan model di `app/models/registry.go`:

```go
var ModelRegistry = []interface{}{
    &User{},
    &Post{}, // tambahkan model baru di sini
}
```

Jalankan migrasi:

```bash
go run main.go migrate
```

### Driver yang Didukung

| Driver | `DB_DRIVER` |
|---|---|
| MySQL | `mysql` |
| PostgreSQL | `postgres` |
| SQLite | `sqlite` |

---

## 💾 Cache

Gunakan interface `contracts.Cache` melalui dependency injection:

```go
import "gons/app/contracts"

type MyService struct {
    Cache contracts.Cache `container:"type"`
}
```

Driver yang tersedia: `memory` (default), `redis` (segera hadir).

---

## 📬 Queue

Gunakan interface `contracts.Queue` melalui dependency injection:

```go
import "gons/app/contracts"

type MyService struct {
    Queue contracts.Queue `container:"type"`
}
```

Driver yang tersedia: `sync`, `goroutine`, `redis` (segera hadir).

---

## 🗃️ Storage

Gunakan interface `contracts.Storage` melalui dependency injection:

```go
import "gons/app/contracts"

type MyService struct {
    Storage contracts.Storage `container:"type"`
}
```

### Konfigurasi S3 / Supabase / MinIO

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

## 🔐 Middleware

### AuthGuard

Middleware `AuthGuard` memeriksa keberadaan header `Authorization`. Tambahkan ke rute manapun:

```go
protected := app.Group("/api")
protected.Use(middlewares.AuthGuard())
```

### Menambah Middleware Kustom

Buat file baru di `app/http/middlewares/` dan implementasikan `fiber.Handler`.

---

## 🎨 Templating (HTML + Vite)

Template HTML disimpan di `resources/views/`. Gunakan helper `vite` di template untuk menyertakan aset yang dikelola Vite:

```html
<!-- resources/views/layouts/main.html -->
<head>
    {{ vite "resources/js/app.js" }}
</head>
```

Build aset frontend:

```bash
npm run build
```

---

## 🛠️ Teknologi yang Digunakan

- [Go](https://go.dev/) — bahasa pemrograman utama
- [Fiber v3](https://gofiber.io/) — web framework
- [GORM](https://gorm.io/) — ORM
- [Golobby Container](https://github.com/golobby/container) — dependency injection
- [godotenv](https://github.com/joho/godotenv) — manajemen file .env
- [go-playground/validator](https://github.com/go-playground/validator) — validasi input
- [Vite](https://vitejs.dev/) — build tool frontend
- [Air](https://github.com/air-verse/air) — hot reload

---

## 📄 Lisensi

Proyek ini dilisensikan di bawah [MIT License](LICENSE).
