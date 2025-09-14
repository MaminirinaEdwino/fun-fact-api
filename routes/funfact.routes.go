package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/fun-fact-api/controller"
)

func Router(mux *http.ServeMux){
	mux.HandleFunc("POST /", controller.Post)
	mux.HandleFunc("GET /", controller.GetAll)
	mux.HandleFunc("GET /{id}", controller.GetById)
	mux.HandleFunc("PUT /{id}", controller.Put)
	mux.HandleFunc("DELETE /{id}", controller.Delete)
}