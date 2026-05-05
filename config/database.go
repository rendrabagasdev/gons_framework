package config

import (
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"
	"time"

	"github.com/golobby/container/v3"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func RegistererDatabase() {
	err := container.Singleton(func() *gorm.DB {
		driver, host, port, user, pass, name := getDatabaseConfig()

		// Attempt to create database if it doesn't exist
		createDatabaseIfNotExist(driver, host, port, user, pass, name)

		// GORM Logging to database.log
		file, _ := os.OpenFile("database.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		newLogger := logger.New(
			log.New(file, "\r\n", log.LstdFlags),
			logger.Config{
				SlowThreshold:             time.Second,
				LogLevel:                  logger.Info,
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      false,
				Colorful:                  false,
			},
		)

		log.Printf("Gons: connecting to database [%s] on [%s:%s]...\n", name, host, port)

		db, err := gorm.Open(getDialector(driver, host, port, user, pass, name), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			log.Printf("Gons: database connection error: %v\n", err)
			slog.Error("Gons: database connection error: " + err.Error())
		}

		// Verify active database
		var currentDB string
		db.Raw("SELECT current_database()").Scan(&currentDB)
		log.Printf("Gons: connection verified! Active database is: [%s]\n", currentDB)

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
		dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pass, host, port, strings.ToLower(name))
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
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/postgres?sslmode=disable", user, pass, host, port)
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
		// Postgres check if exists (Postgres stores identifiers in lowercase)
		var exists int
		db.Raw("SELECT 1 FROM pg_database WHERE LOWER(datname) = LOWER(?)", name).Scan(&exists)
		if exists == 0 {
			log.Printf("Gons: database [%s] not found, creating now...\n", name)
			db.Exec(fmt.Sprintf("CREATE DATABASE %s", strings.ToLower(name)))
		}
	}
}
