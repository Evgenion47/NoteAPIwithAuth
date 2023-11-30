package middleware

import (
	"API/internal/models"
	"API/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	var note models.CreateNote
	err := json.NewDecoder(r.Body).Decode(&note)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	insertId := repository.CreateNote(note)

	res := models.Response{
		ID:      insertId,
		Message: "Note created successfully",
	}

	json.NewEncoder(w).Encode(res)
}
