package main

import (
	"encoding/json"
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
	
	if login.Username != "Admin" || login.Password != "admin" {
		http.Error(w, "That is not a valid login. Please try again.", 500)
		return
	}

	fmt.Fprintf(w, "Welcome back, %v!", login.Username)
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

