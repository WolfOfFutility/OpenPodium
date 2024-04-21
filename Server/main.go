package main

import (
	"fmt"
	"net/http"
)

// Handle CORS - This should be locked down in the future
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func testSecretVault() {
	writeSecretToUserVault("Admin", map[string]any{"SecondToken": "Token1234"})
	readSecretsFromUserVault("Admin")
	//removeSecretFromUserVault("Admin", "SecondToken")
}

// Main function - start server
func main() {
	// testSecretVault()

	// runVault()

	// Server start message
	listenPort := 4000
	fmt.Printf("Running server on port %v\n", listenPort)
	
	// Endpoint Handlers
	http.HandleFunc("/login", podiumLogin)

	// Endpoints - Secret Management
	http.HandleFunc("/secrets", getUserSecrets)
	http.HandleFunc("/secrets/new", addUserSecret)
	http.HandleFunc("/secrets/remove", removeUserSecret)

	// http.HandleFunc("/", index)
	// http.HandleFunc("/headers", headers)
	// http.HandleFunc("/hello", hello)
	
	// http.HandleFunc("/azure", azureLogin)
	

	// Start server and listen on the specified port
	http.ListenAndServe(fmt.Sprintf(":%v", listenPort), nil)
}
