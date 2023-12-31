package main

import (
	"API/internal/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on the port 8000...")

	log.Fatal(http.ListenAndServe(":8000", r))
}
