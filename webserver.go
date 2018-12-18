package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"bufio"
)

func createHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	message := map[string]interface{}{
		"settings": map[string]int{
			"number_of_shards":   1,
			"number_of_replicas": 0,
		},
	}
	bytesRepresentation, err := json.Marshal(message)
	var url string
	url = "http://localhost:9200/" + params["id"]
	req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(bytesRepresentation))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		// Handle error
		panic(err)
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	fmt.Fprintf(w, "index %s created", params["id"])
}

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var url string
        url = "http://localhost:9200/" + params["id"]
	req, err := http.NewRequest(http.MethodDelete, url, nil)
        req.Header.Set("Content-Type", "application/json")
        if err != nil {
                // Handle error
                panic(err)
        }
        client := &http.Client{}
        resp, err := client.Do(req)
        defer resp.Body.Close()
	fmt.Fprintf(w, "%s deleted", params["id"])
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://localhost:9200/_cat/indices"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	bs := string(content)
	scanner := bufio.NewScanner(strings.NewReader(bs))
	var line string
	for scanner.Scan() {
		line = scanner.Text()
		index := strings.Split(line," ")
		fmt.Println(index[2])
	}
	fmt.Fprintf(w, string(content))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/get", getHandler).Methods("GET")
	router.HandleFunc("/test/{id}", createHandler).Methods("GET")
	router.HandleFunc("/test/{id}", deleteHandler).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", router))
}
