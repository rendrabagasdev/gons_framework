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
		db, err := gorm.Open(getDriver(), &gorm.Config{})
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

func getDriver() gorm.Dialector {
	driver := GetEnv("DB_DRIVER", "mysql")
	host := GetEnv("DB_HOST", "localhost")
	port := GetEnv("DB_PORT", "3306")
	user := GetEnv("DB_USER", "root")
	pass := GetEnv("DB_PASS", "")
	name := GetEnv("DB_NAME", "")

	switch driver {
	case "mysql":
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
		return mysql.Open(dsn)
	case "postgres":
		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, pass, name, port)
		return postgres.Open(dsn)
	case "sqlite":
		return sqlite.Open(name + ".db")
	default:
		slog.Error("Gons: database driver [" + driver + "] is not supported")
		return nil
	}
}
