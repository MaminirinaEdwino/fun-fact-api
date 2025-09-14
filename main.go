package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MaminirinaEdwino/fun-fact-api/routes"
	"github.com/rs/cors"
)

func main() {
	mux := http.NewServeMux()
	cors := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowCredentials: true,
	})
	routes.Router(mux)
	fmt.Println("Server started on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", cors.Handler(mux)))
}
