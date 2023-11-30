package repository

import (
	"API/internal/db"
	"API/internal/models"
	"log"
)

func GetAll() (models.Notes, error) {
	db := db.CreateConn()
	defer db.Close()

	var notes models.Notes

	query := `select * from notes`
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var note models.Note
		err = rows.Scan(&note.ID, &note.OwnerID, &note.Title, &note.Description, &note.Completion, &note.CreatedAt)
		if err != nil {
			log.Fatalf("Unable to scan the row. %v", err)
		}
		notes.Notes = append(notes.Notes, note)
	}

	return notes, err
}
