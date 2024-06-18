package persistence

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
		host:     "host.docker.internal",  //os.Getenv("MYSQL_HOST"),
		port:     "3306",                  // os.Getenv("MYSQL_PORT"),
		username: "goferProcessor",        // os.Getenv("MYSQL_USER"),
		password: "_GoProcessor#2024_",    // os.Getenv("MYSQL_PASSWORD"),
		database: "transaction-processor", // os.Getenv("MYSQL_DATABASE_NAME"),
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
	return db, nil
}
