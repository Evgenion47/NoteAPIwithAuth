package repository

import (
	"API/internal/db"
	"API/internal/models"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func Register(credentials models.AuthInfo) error {
	dbconn := db.CreateConn()
	query := `insert into users (userid, password) values ($1,$2)`
	var id int64
	password, _ := bcrypt.GenerateFromPassword([]byte(credentials.Password), 14)
	err := dbconn.QueryRow(query, credentials.UserId, password)
	if err != nil {
		return fmt.Errorf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return nil
}
