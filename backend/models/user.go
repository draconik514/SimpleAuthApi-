package models 

import (
	"backend/config"
	"database/sql"

	"time"
	"errors"
)

type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"-"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateUser(name, email, password string) error {
	db := config.GetDB()

	query := "INSERT INTO (name, email, password) VALUES (?, ?, ?)"	
	_, err := db.Exec(query, name, email, password)

	return err
}

func GetUserByEmail(email string) (*User, error){
	db := config.GetDB()

	user := &User{}
	query := "SELECT id, name, email, password, created_at FROM users WHERE email = ?"

	err := db.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows{
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return user, nil
}

func GetUserByID(id int) (*User, error){
	db := config.GetDB()

	user := &User{}
	query := "SELECT name, email, password, created_at FROM users WHERE id = ?"

	err := db.QueryRow(query, id).Scan(
		&user.ID,
		&user.Name, 
		&user.Email,
		&user.Password,
		&user.CreatedAt, 
	)

	if err != nil {
		return nil, err
	}

	return user, nil
}