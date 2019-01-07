package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	test = prometheus.NewCounterVec(prometheus.CounterOpts{Name: "test", Help: "test metric"},[]string{"endpoint"},)
)

func testHandler(w http.ResponseWriter, r *http.Request) {
	test.WithLabelValues("/test").Inc()
	fmt.Fprintf(w, "test")
}

func init() {
	prometheus.MustRegister(test)
}

func fooHandler(w http.ResponseWriter, r *http.Request) {
        test.WithLabelValues("/foo").Inc()
        fmt.Fprintf(w, "bar")
}


func main() {
	//This section will start the HTTP server and expose
	//any metrics on the /metrics endpoint.
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/test", testHandler)
	http.HandleFunc("/foo", fooHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
