package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handle CORS - This should be locked down in the future
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func testSecretVault() {
	// enableUserAuth()
	// createVaultUser(Login{Username: "Admin", Password: "admin"}, Login{Username: "root", Password: "root"})
	
	auth, err := loginVaultUser(Login{Username: "Admin", Password: "admin"})
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(auth)
}

// Main function - start server
func main() {
	testSecretVault()

	// Server start message
	listenPort := 4000
	fmt.Printf("Running server on port %v\n", listenPort)
	
	// Endpoint Handlers
	// http.HandleFunc("/login", podiumLogin)

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
