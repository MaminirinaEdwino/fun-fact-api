package routes

import (
	"net/http"

	"github.com/MaminirinaEdwino/fun-fact-api/controller"
)

func Router(mux *http.ServeMux){
	// mux.Handle("GET /", http.FileServer(http.Dir("./Template")))
	mux.HandleFunc("GET /", controller.MainController)
	mux.HandleFunc("POST /funfact", controller.Post)
	mux.HandleFunc("GET /funfact", controller.GetAll)
	mux.HandleFunc("GET /funfact/{id}", controller.GetById)
	mux.HandleFunc("PUT /funfact/{id}", controller.Put)
	mux.HandleFunc("DELETE /funfact/{id}", controller.Delete)
}