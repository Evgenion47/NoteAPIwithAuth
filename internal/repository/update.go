package repository

import (
	"API/internal/db"
	"API/internal/models"
	"fmt"
	"log"
)

func UpdateNote(id models.NoteId, note models.UpdateNote) int64 {
	db := db.CreateConn()
	defer db.Close()
	query := `update notes set title=$2, description=$3, completion=$4 where noteid=$1`
	res, err := db.Exec(query, id.ID, note.Title, note.Description, note.Completion)

	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	// check how many rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("Error while checking the affected rows. %v", err)
	}

	fmt.Printf("Total rows/record affected %v", rowsAffected)

	return rowsAffected
}
