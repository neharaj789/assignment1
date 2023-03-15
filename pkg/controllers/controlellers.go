package controllers

import (
	"encoding/json"
	"net/http"

	"assignment1.1/pkg/data"
	"github.com/gorilla/mux"
)

//var m sync.Mutex

func GetViewcountbyID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	par := mux.Vars(r)

	//var it int
	//m.Lock()
	for _, item := range data.Videoplat {
		if item.Id == par["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	//m.Unlock()

}

func IncrViewCountbyID(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//m.Lock()
	for i, _ := range data.Videoplat {
		if data.Videoplat[i].Id == params["id"] {

			data.Videoplat[i].Viewcount++

			json.NewEncoder(w).Encode(data.Videoplat[i])
			return

		}
	}
	//m.Unlock()
}
