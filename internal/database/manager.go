package database

import (
	"fmt"
	"gons/pkg/env"
	"log/slog"
	"os"
	"strconv"
	"strings"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewGormConnection initializes and returns a GORM database connection.
func NewGormConnection() *gorm.DB {
	driver := env.Get("DB_CONNECTION", "sqlite")
	host := env.Get("DB_HOST", "127.0.0.1")
	port := env.Get("DB_PORT", "3306")
	database := env.Get("DB_DATABASE", "gons.db")
	username := env.Get("DB_USERNAME", "root")
	password := env.Get("DB_PASSWORD", "")

	var dialector gorm.Dialector

	switch strings.ToLower(driver) {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, port, database)
		dialector = mysql.Open(dsn)
	case "postgres", "postgresql":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC", host, username, password, database, port)
		dialector = postgres.Open(dsn)
	case "sqlite":
		dialector = sqlite.Open(database)
	default:
		slog.Error("Unsupported database driver", "driver", driver)
		os.Exit(1)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	sqlDB, err := db.DB()
	if err != nil {
		slog.Error("Failed to get generic database object", "error", err)
		os.Exit(1)
	}

	// Connection Pooling
	maxIdleConns, _ := strconv.Atoi(env.Get("DB_MAX_IDLE_CONNS", "10"))
	maxOpenConns, _ := strconv.Atoi(env.Get("DB_MAX_OPEN_CONNS", "100"))
	connMaxLifetimeStr := env.Get("DB_CONN_MAX_LIFETIME", "1h")
	connMaxLifetime, err := time.ParseDuration(connMaxLifetimeStr)
	if err != nil {
		connMaxLifetime = time.Hour
	}

	sqlDB.SetMaxIdleConns(maxIdleConns)
	sqlDB.SetMaxOpenConns(maxOpenConns)
	sqlDB.SetConnMaxLifetime(connMaxLifetime)

	slog.Info("Database connected successfully", "driver", driver, "database", database)

	return db
}
