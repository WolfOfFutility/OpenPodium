package main

import (
	// "fmt"
	"net/http"
	"github.com/WolfOfFutility/OpenPodium/Server/handlers"
	"github.com/WolfOfFutility/OpenPodium/Server/testers"
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
	testSecretVault()
	handlers.testHandlerModule()
	testers.testHandlerModule()

	// runVault()

	// Server start message
	// listenPort := 4000
	//fmt.Printf("Running server on port %v\n", listenPort)
	
	// Endpoint Handlers
	// http.HandleFunc("/", index)
	// http.HandleFunc("/headers", headers)
	// http.HandleFunc("/hello", hello)
	// http.HandleFunc("/login", podiumLogin)
	// http.HandleFunc("/azure", azureLogin)
	// http.HandleFunc("/secrets/new", addUserSecret)
	// http.HandleFunc("/secrets", getUserSecrets)

	// Start server and listen on the specified port
	//http.ListenAndServe(fmt.Sprintf(":%v", listenPort), nil)
}
