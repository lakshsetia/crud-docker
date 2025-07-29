package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/lakshsetia/crud-docker/internal/config"
	"github.com/lakshsetia/crud-docker/internal/types"
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
func (p *Postgresql) GetUsers() ([]types.User, error) {
	rows, err := p.DB.Query("SELECT id, name, email, age FROM users")
	if err != nil {
		return nil, err
	}
	users := make([]types.User, 0)
	for rows.Next() {
		var user types.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Email, &user.Age); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
func (p *Postgresql) CreateUser(user types.User) error {
	if _, err := p.DB.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", user.Name, user.Email, user.Age); err != nil {
		return err
	}
	return nil
}
func (p *Postgresql) GetUserById(id int) (types.User, error) {
	var user types.User
	err := p.DB.QueryRow("SELECT id, name, email, age FROM users WHERE id=$1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	if err == sql.ErrNoRows {
		return user, fmt.Errorf("user not found with id=%v", id)
	}
	if err != nil {
		return user, err
	}
	return user, nil
}
func (p *Postgresql) UpdateUserById(id int, newUser types.User) error {
	if _, err := p.DB.Exec("UPDATE users SET name=$1, email=$2, age=$3 WHERE id=$4", newUser.Name, newUser.Email, newUser.Age, id); err != nil {
		return err
	}
	return nil
}
func (p *Postgresql) DeleteUserById(id int) error {
	if _, err := p.DB.Exec("DELETE FROM users WHERE id=$1", id); err != nil {
		return err
	}
	return nil
}