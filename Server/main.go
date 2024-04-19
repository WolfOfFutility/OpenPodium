package main

import (
	"fmt"
	"net/http"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func main() {
	listenPort := 4000

	fmt.Printf("Running server on port %v", listenPort)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/hello", hello)
	// http.HandleFunc("/azure", handlers.azureLogin)

	http.ListenAndServe(fmt.Sprintf(":%v", listenPort), nil)
}
