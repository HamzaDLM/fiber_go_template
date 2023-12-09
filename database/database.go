package database

import (
	"strings"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

type DatabaseConfig struct {
	Driver   string
	Host     string
	Username string
	Password string
	Port     string
	Database string
}

type Database struct {
	*gorm.DB
}

func New(config *DatabaseConfig, logger *zap.Logger) (*Database, error) {
	var db *gorm.DB
	var err error
	switch strings.ToLower(config.Driver) {
	case "mysql":
		dsn := config.Username + ":" + config.Password + "@tcp(" + config.Host + ":" + config.Port + ")/" + config.Database + "?charset=utf8mb4&collation=utf8mb4_general_ci&parseTime=True&loc=UTC"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		break
	case "postgresql", "postgres":
		dsn := "user=" + config.Username + " password=" + config.Password + " dbname=" + config.Database + " host=" + config.Host + " port=" + config.Port + " TimeZone=UTC"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		break
	case "sqlserver", "mssql":
		dsn := "sqlserver://" + config.Username + ":" + config.Password + "@" + config.Host + ":" + config.Port + "?database=" + config.Database
		db, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
		break
	case "sqlite", "sqlite3":
		db, err = gorm.Open(sqlite.Open(config.Host), &gorm.Config{})
		break
	default: 
		logger.Fatal("Database configured dialect/driver not recognized")
	}
	return &Database{db}, err
}
