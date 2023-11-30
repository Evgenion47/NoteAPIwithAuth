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
	query := `select password from users where userid=$1`
	err := dbconn.QueryRow(query, credentials.UserId).Scan(&pass)
	if err != nil {
		log.Printf("Unable to execute the query. %v", err)
	}
	if err = bcrypt.CompareHashAndPassword([]byte(pass), []byte(credentials.Password)); err != nil {
		log.Printf("%v", err)
		return false
	}
	return true
}
