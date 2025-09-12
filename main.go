package main

import (
	// "database/sql"
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	Router(mux)
	fmt.Println("Server started on localhost:8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
