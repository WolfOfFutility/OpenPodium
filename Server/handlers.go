package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/google/uuid"
	"log"
)

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

// func index(w http.ResponseWriter, req *http.Request) {
// 	fmt.Fprintf(w, "Main Page")
// }

func hello(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	fmt.Fprintf(w, "hello\n")
}

// Baseline login to the application, rudimentary right now
// * TO DO:
// ** Store credentials in an encrypted format, check for matches
// ** Needs to be synced to a Secure Store, give user access at a certain level to see their saved secrets on their account
// ** Basic SSO?
func podiumLogin(w http.ResponseWriter, req *http.Request) {
	enableCors(&w)

	var login Login

	err := json.NewDecoder(req.Body).Decode(&login)
	
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	
	// Check for logins
	if login.Username != "Admin" || login.Password != "admin" {
		http.Error(w, "That is not a valid login. Please try again.", 500)
		return
	}

	// Set User SSOToken
	SSOToken := uuid.New()
	writeSecretToUserVault(login.Username, map[string]any{"SSOToken": SSOToken})

	json.NewEncoder(w).Encode(map[string]any{"Message": fmt.Sprintf("Welcome back, %v!", login.Username), "SSOToken": SSOToken})

	//fmt.Fprintf(w, "Welcome back, %v!", login.Username)
}

func addUserSecret(w http.ResponseWriter, req *http.Request) {
	enableCors(&w) 
	
	userName := "Admin"
	testSecret := map[string]any{
		"DevOps_Secret": "devopssecret1",
	}

	writeSecretToUserVault(userName, testSecret)
}

func getUserSecrets(w http.ResponseWriter, req *http.Request) {
	enableCors(&w) 
	
	userName := "Admin"

	userSecrets, err := readSecretsFromUserVault(userName)
	if err != nil {
		log.Fatal(err)
	}

	// Pull SSOToken out of the secret store that is returned to the user
	showableSecrets := make(map[string]any)
	for key, value := range userSecrets {
		if key != "SSOToken" {
			showableSecrets[key] = value
		}
	}

	// If no SSOToken is sent, return an error
	if userSecrets["SSOToken"] == nil {
		json.NewEncoder(w).Encode(map[string]any{"Code": 400, "Message": "Invalid SSOToken, please login."})
	} else{
		json.NewEncoder(w).Encode(showableSecrets)
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

