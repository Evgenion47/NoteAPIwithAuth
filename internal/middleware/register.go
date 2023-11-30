package middleware

import (
	"API/internal/models"
	"API/internal/repository"
	"encoding/json"
	"log"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var credentials models.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Printf("Unable to decode the request body.  %v", err)
	}
	err = repository.Register(credentials)
	if err != nil {
		log.Printf("%v", err)
	}

}
