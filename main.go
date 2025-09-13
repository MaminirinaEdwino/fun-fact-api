package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/MaminirinaEdwino/fun-fact-api/routes"
)

func main() {
	mux := http.NewServeMux()
	routes.Router(mux)
	fmt.Println("Server started on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
