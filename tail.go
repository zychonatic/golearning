package main

import (
	"github.com/hpcloud/tail"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"log"
)

func tailHandler (w http.ResponseWriter, r *http.Request){
	t, err := tail.TailFile("./test.log", tail.Config{Follow: false})
	if err != nil {
                // Handle error
                panic(err)
        }
	fmt.Fprintf(w, "<html><head><meta http-equiv=\"refresh\" content=\"1\" /></head><body>")
	for line := range t.Lines {
		fmt.Fprintf(w, line.Text + "<br>")
	}
}

func main (){
	router := mux.NewRouter()
        router.HandleFunc("/tail", tailHandler)
	log.Fatal(http.ListenAndServe(":8080", router))
}
