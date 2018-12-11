package main

import (
	"fmt"
	"log"
	"net/http"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I will create %s!", r.URL.Path[len("/create/"):])
}

func deleteHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "this will delete %s", r.URL.Path[len("/delete/"):])
}

func main() {
	http.HandleFunc("/create/", createHandler)
	http.HandleFunc("/delete/", deleteHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
