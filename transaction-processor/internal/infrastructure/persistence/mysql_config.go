package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
	"time"
)

type DBConfig struct {
	host     string
	port     string
	username string
	password string
	database string
}

func NewDBConfig() *DBConfig {
	return &DBConfig{
		host:     os.Getenv("MYSQL_HOST"),
		port:     os.Getenv("MYSQL_PORT"),
		username: os.Getenv("MYSQL_USER"),
		password: os.Getenv("MYSQL_PASSWORD"),
		database: os.Getenv("MYSQL_DATABASE_NAME"),
	}
}

func (config *DBConfig) ConnectToDB() (*sql.DB, error) {
	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=true",
		config.username,
		config.password,
		config.host,
		config.port,
		config.database,
	)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}
