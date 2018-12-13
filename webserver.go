package main

import (
	"fmt"
	"log"
	"net/http"
	"io/ioutil"
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

func getHandler(w http.ResponseWriter, r *http.Request){
	url := "http://localhost:9200/_cat/indices"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	fmt.Fprintf(w, string(content))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get", getHandler).Methods("GET")
	router.HandleFunc("/test/{id}", createHandler).Methods("POST")
	router.HandleFunc("/test/{id}", deleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
