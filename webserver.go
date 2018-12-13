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
	var url string
	url = "http://localhost:9200/" + params["id"]
	req, err := http.NewRequest(http.MethodPut, url, nil)
	if err != nil {
	// Handle error
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Fprintf(w, "index %s created", params["id"])
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
	router.HandleFunc("/test/{id}", createHandler).Methods("GET")
	router.HandleFunc("/test/{id}", deleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
