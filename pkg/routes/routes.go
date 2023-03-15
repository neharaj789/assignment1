package routes

import (
	"assignment1.1/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookroutes = func(r *mux.Router) {

	go r.HandleFunc("/videos/{id}", controllers.IncrViewCountbyID).Methods("POST")
	go r.HandleFunc("/videos/viewcount/{id}", controllers.GetViewcountbyID).Methods("GET")
}
