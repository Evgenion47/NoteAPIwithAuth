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

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	noteId := models.NoteId{ID: int64(id), OwnerID: r.Header.Get("name")}
	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	deletedRows := repository.DeleteNote(noteId)
	res := models.Response{}
	if deletedRows == 0 {
		res.ID = -1
		res.Message = "Record not found or permission denied"
		w.WriteHeader(404)
	} else {
		msg := fmt.Sprintf("Note updated successfully. Total rows/record affected %v", deletedRows)
		res.ID = int64(id)
		res.Message = msg
	}
	json.NewEncoder(w).Encode(res)
}
