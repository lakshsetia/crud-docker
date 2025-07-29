package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/lakshsetia/crud-docker/internal/config"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	DB *sql.DB
}

func New(config *config.Config) (*Postgresql, error) {
	user, passwd, dbname, host, port := config.Database.Postgresql.User, config.Database.Postgresql.Password, config.Database.Postgresql.DBName, config.Database.Postgresql.Host, config.Database.Postgresql.Port
	if user == "" || passwd == "" || dbname == "" || host == "" || port == "" {
		return nil, fmt.Errorf("database connection parameters not specified")
	}
	connectionStr := fmt.Sprintf("user=%v password=%v dbname=%v host=%v port=%v sslmode=disable", user, passwd, dbname, host, port)
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	email VARCHAR(200) NOT NULL,
	age INTEGER NOT NULL CHECK (age >= 0)
	)
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to create database table: %w", err)
	}
	return &Postgresql{
		DB: db,
	}, nil
}