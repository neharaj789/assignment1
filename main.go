package main

import (
	"fmt"
	"log"
	"net/http"

	"assignment1.1/pkg/routes"
	"assignment1.1/pkg/video"
	"github.com/gorilla/mux"
)

var videoplat []video.Video

func main() {

	//data.Myfunc()
	r := mux.NewRouter()
	routes.RegisterBookroutes(r)
	http.Handle("/", r)

	fmt.Println("Starting server at port 9010\n")

	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
