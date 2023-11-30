package repository

import (
	"API/internal/db"
	"API/internal/models"
	"fmt"
	"log"
)

func CreateNote(note models.CreateNote) int64 {
	db := db.CreateConn()
	query := `insert into notes (title, description) values ($1,$2) returning noteid`
	var id int64
	err := db.QueryRow(query, note.Title, note.Description).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Printf("Inserted a single record %v", id)
	return id
}
