package middleware

import (
	"API/internal/models"
	"API/internal/repository"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func UpdateNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	noteId := models.NoteId{ID: int64(id)}
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	var note models.UpdateNote

	err = json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := repository.UpdateNote(noteId, note)

	// format the message string
	msg := fmt.Sprintf("Note updated successfully. Total rows/record affected %v", updatedRows)

	// format the response message
	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}
