package config

import (
	"fmt"
	"log/slog"

	"github.com/golobby/container/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func RegistererDatabase() {
	err := container.Singleton(func() *gorm.DB {
		driver, host, port, user, pass, name := getDatabaseConfig()

		// Attempt to create database if it doesn't exist
		createDatabaseIfNotExist(driver, host, port, user, pass, name)

		slog.Info(fmt.Sprintf("Gons: connecting to database [%s] on [%s:%s]...", name, host, port))

		db, err := gorm.Open(getDialector(driver, host, port, user, pass, name), &gorm.Config{})
		if err != nil {
			slog.Error("Gons: database connection error: " + err.Error())
		}
		return db
	})
	if err != nil {
		slog.Error("Gons: database register error: " + err.Error())
	}
}

func init() {
	RegisterConfig(func() error {
		RegistererDatabase()
		return nil
	})
}

func getDatabaseConfig() (driver, host, port, user, pass, name string) {
	driver = GetEnv("DB_DRIVER", "mysql")
	host = GetEnv("DB_HOST", "localhost")
	port = GetEnv("DB_PORT", "3306")
	user = GetEnv("DB_USER", "root")
	pass = GetEnv("DB_PASS", "")
	name = GetEnv("DB_NAME", "")
	return
}

func getDialector(driver, host, port, user, pass, name string) gorm.Dialector {
	switch driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
		return mysql.Open(dsn)
	case "postgres", "postgresql", "psql", "postgree":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)
		return postgres.Open(dsn)
	case "sqlite":
		return sqlite.Open(name + ".db")
	default:
		slog.Error("Gons: database driver [" + driver + "] is not supported")
		return nil
	}
}

func createDatabaseIfNotExist(driver, host, port, user, pass, name string) {
	if name == "" || driver == "sqlite" {
		return
	}

	var dsn string
	var dialector gorm.Dialector

	if driver == "mysql" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", user, pass, host, port)
		dialector = mysql.Open(dsn)
	} else if driver == "postgres" || driver == "postgresql" || driver == "psql" || driver == "postgree" {
		// Connect to default 'postgres' database to create the new one
		dsn = fmt.Sprintf("host=%s user=%s password=%s dbname=postgres port=%s sslmode=disable", host, user, pass, port)
		dialector = postgres.Open(dsn)
	} else {
		return
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	if driver == "mysql" {
		db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", name))
	} else {
		// Postgres check if exists
		var exists int
		db.Raw("SELECT 1 FROM pg_database WHERE datname = ?", name).Scan(&exists)
		if exists == 0 {
			db.Exec(fmt.Sprintf("CREATE DATABASE %s", name))
		}
	}
}
