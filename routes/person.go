package routes

import (
	"latihan_golang/controllers/personcontroller"

	"github.com/gorilla/mux"
)

func PersonRoutes(r *mux.Router) {
	router := r.PathPrefix("/person").Subrouter()

	router.HandleFunc("", personcontroller.Index).Methods("GET")
	router.HandleFunc("", personcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", personcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", personcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", personcontroller.Delete).Methods("DELETE")

}
