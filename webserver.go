package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Hi there, I will create %s!", params["id"])
}

func deleteHandler(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	fmt.Fprintf(w, "this will delete %s", params["id"])
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/test/{id}", createHandler).Methods("POST")
	router.HandleFunc("/test/{id}", deleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
