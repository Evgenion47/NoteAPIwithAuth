package repository

import (
	"API/internal/db"
	"API/internal/models"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func Authenticate(credentials models.AuthInfo) bool {
	dbconn := db.CreateConn()
	var pass string
	query := `select userid from users where userid=$1 and password=$2)`
	err := dbconn.QueryRow(query, credentials.UserId, credentials.Password).Scan(&pass)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}
	if bcrypt.CompareHashAndPassword([]byte(pass), []byte(credentials.Password)) != nil {
		return false
	}
	return true
}
