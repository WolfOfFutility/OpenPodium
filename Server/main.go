package main

import (
	"fmt"
	"net/http"
)

// Handle CORS - This should be locked down in the future
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

// Main function - start server
func main() {
	listenPort := 4000

	// Server start message
	fmt.Printf("Running server on port %v", listenPort)
	
	// Endpoint Handlers
	// http.HandleFunc("/", index)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/login", podiumLogin)
	http.HandleFunc("/azure", azureLogin)

	// Start server and listen on the specified port
	http.ListenAndServe(fmt.Sprintf(":%v", listenPort), nil)
}
