package router

import (
	"API/internal/middleware"
	"github.com/gorilla/mux"
	"net/http"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/", middleware.GetAll).Methods("GET", "OPTIONS")
	router.HandleFunc("/authenticate/", middleware.Authenticate).Methods("POST", "OPTIONS")
	router.HandleFunc("/register/", middleware.Register).Methods("POST", "OPTIONS")
	router.Handle("/api/v1/{id}", middleware.AuthMiddleware(http.HandlerFunc(middleware.CreateNote))).Methods("POST", "OPTIONS")
	router.Handle("/api/v1/{id}", middleware.AuthMiddleware(http.HandlerFunc(middleware.UpdateNote))).Methods("PUT", "OPTIONS")
	router.Handle("/api/v1/", middleware.AuthMiddleware(http.HandlerFunc(middleware.CreateNote))).Methods("POST", "OPTIONS")
	router.Handle("/logout/", middleware.AuthMiddleware(http.HandlerFunc(middleware.Logout))).Methods("GET", "OPTIONS")

	return router
}
