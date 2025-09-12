package main

import "net/http"

func Router(mux *http.ServeMux){
	mux.HandleFunc("POST /", GetAll)
	mux.HandleFunc("GET /", GetAll)
	mux.HandleFunc("GET /{id}", GetById)
}