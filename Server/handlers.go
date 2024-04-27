package main

import (
	// "encoding/json"
	"fmt"
	"net/http"
	// "github.com/google/uuid"
	// "log"
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

