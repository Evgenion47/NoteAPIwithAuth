package middleware

import (
	"API/internal/models"
	"API/internal/repository"
	"API/internal/utils"
	"encoding/json"
	"log"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	var credentials models.AuthInfo
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Fatalf("Unable to decode the request body.  %v", err)
	}
	if repository.Authenticate(credentials) {
		token, err := utils.GetToken(credentials.UserId)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating JWT token: " + err.Error()))
		} else {
			w.Header().Set("Authorization", "Bearer "+token)
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Token: " + token))
		}
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Name and password do not match"))
		return
	}
}
