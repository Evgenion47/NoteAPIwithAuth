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
	query := `update notes set title=$3, description=$4, completion=$5 where noteid=$1 and ownerid=$2`
	res, err := db.Exec(query, id.ID, id.OwnerID, note.Title, note.Description, note.Completion)

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
