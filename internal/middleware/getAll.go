package middleware

import (
	"API/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	notes, err := repository.GetAll()
	if err != nil {
		log.Println("occurred error while get all notes")
	}
	json.NewEncoder(w).Encode(notes)
}
