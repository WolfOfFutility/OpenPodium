package main

import (
	"fmt"
	// "log"
	"net/http"
)

// Handle CORS - This should be locked down in the future
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func initSecretVault() {
	enableUserAuth()
	createVaultUser(Login{Username: "Admin", Password: "admin"}, Login{Username: "root", Password: "root"})
}

// Main function - start server
func main() {
	initSecretVault()

	// Server start message
	listenPort := 8080
	fmt.Printf("Running server on port %v\n", listenPort)
	
	// Endpoint Handlers
	http.HandleFunc("/login", podiumLogin)

	// Endpoints - Secret Management
	// http.HandleFunc("/secrets", getUserSecrets)
	// http.HandleFunc("/secrets/new", addUserSecret)
	// http.HandleFunc("/secrets/remove", removeUserSecret)

	// http.HandleFunc("/", index)
	// http.HandleFunc("/headers", headers)
	// http.HandleFunc("/hello", hello)
	
	// http.HandleFunc("/azure", azureLogin)
	

	// Start server and listen on the specified port
	http.ListenAndServe(fmt.Sprintf(":%v", listenPort), nil)
}
