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
	noteId := models.NoteId{ID: int64(id), OwnerID: r.Header.Get("name")}
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	var note models.UpdateNote

	err = json.NewDecoder(r.Body).Decode(&note)

	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}

	updatedRows := repository.UpdateNote(noteId, note)

	res := models.Response{}

	if updatedRows == 0 {
		res.ID = -1
		res.Message = "Record not found or permission denied"
		w.WriteHeader(404)
	} else {
		msg := fmt.Sprintf("Note updated successfully. Total rows/record affected %v", updatedRows)
		res.ID = int64(id)
		res.Message = msg
	}

	json.NewEncoder(w).Encode(res)

}
