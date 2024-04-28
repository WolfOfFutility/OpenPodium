package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	// "github.com/google/uuid"
	"log"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func hello(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "hello\n")
}

// Handle a login from the front end
func podiumLogin(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	// Initialise Request Object Pointer
	var parsedRequest *Login
	
	// Parse Request Body
	err := json.NewDecoder(req.Body).Decode(&parsedRequest)
	if err != nil {
		log.Println("Decode Error: ", err.Error())
	}

	// Use Request body to pass credentials to login, pass resulting auth object
	auth, errCode, err := loginVaultUser(Login{Username: parsedRequest.Username, Password: parsedRequest.Password})
	if err != nil { // ** May need to work on the logic here, give users a more intuitive response
		log.Println("Login Error: ", errCode, " - ", err.Error())

		errorMessage := "Login Failure."

		if errCode == 400 {
			errorMessage = "Invalid username or password. Please try again."
		}
		
		json.NewEncoder(w).Encode(map[string]any{"Error": errorMessage, "ErrorCode": errCode})
	} else {
		log.Println("Return Auth Values...")
		json.NewEncoder(w).Encode(auth)
	}
}

func removeUserSecret(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)
}

// Placeholder to login to Azure / Microsoft-based authentication
func azureLogin(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "Azure Login Handler")
}

// Placeholder to login to AWS authentication
func AWSLogin(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "AWS Login Handler")
}

// Placeholder to login to Jenkins authentication
func jenkinsLogin(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "Jenkins Login Handler")
}

// Placeholder to login to Jira authentication
func jiraLogin(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "Jira Login Handler")
}

// Placeholder to login to Git authentication
func gitLogin(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "Git Login Handler")
}

