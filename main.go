package main

import (
	"github.com/gorilla/mux"
	"net/http"
)
func newRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/coworkingspace", getCoworkingSpacesHandler).Methods("GET")
	r.HandleFunc("/coworkingspace", createCoworkingSpaceHandler).Methods("POST")
	staticFileDirectory := http.Dir("./assets")
	staticFileHandler := http.StripPrefix("/assets/", http.FileServer(staticFileDirectory))
	r.PathPrefix("/assets").Handler(staticFileHandler).Methods("GET")
	return r
}

func main() {
	r := newRouter()
	http.ListenAndServe(":8080", r)
}
