package repository

import (
	"API/internal/db"
	"API/internal/models"
	"fmt"
	"log"
)

func DeleteNote(id models.NoteId) int64 {
	db := db.CreateConn()
	defer db.Close()
	query := `delete from notes where noteid=$1`
	res, err := db.Exec(query, id.ID)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}
	fmt.Printf("Total rows/record affected %v", rowsAffected)
	return rowsAffected
}
