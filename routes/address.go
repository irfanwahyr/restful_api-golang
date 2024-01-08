package routes

import (
	"latihan_golang/controllers/addresscontroller"
	"github.com/gorilla/mux"
)

func AddressRoutes(r *mux.Router) {
	router := r.PathPrefix("/address").Subrouter()

	router.HandleFunc("", addresscontroller.Index).Methods("GET")
	router.HandleFunc("", addresscontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", addresscontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", addresscontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", addresscontroller.Delete).Methods("DELETE")

}
