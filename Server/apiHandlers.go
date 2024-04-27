package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Structure an API Request for easy access
func makeAPIRequest(method string, URI string, headers http.Header, body map[string]any) (map[string]any, error) {
	// Parse the body that was put in the input
	requestBody, err := json.Marshal(body)
	if err != nil {
		log.Println("Fatal Error on Body Parsing: ", err.Error());
		return nil, err
	}

	// Set request body buffer and initialise http client
	bufferedRequestBody := bytes.NewBuffer(requestBody)
	client := http.Client{}

	// Create the request object
	req, err := http.NewRequest(method, URI, bufferedRequestBody)
	if err != nil {
		log.Println("Fatal Error on Request Creation: ", err.Error());
		return nil, err
	}

	// Add each of the headers specified
	req.Header = headers

	// Execute the Request
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Fatal Error on API Execution: ", err.Error());
		return nil, err
	}

	// Add a defer for the cleanup of the response body afterwards
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Fatal Error on Response Read: ", err.Error());
		return nil, err
	}

	// Parse the response body from JSON format into a Map
	// If there is an empty response, just return an empty map
	responseMap := make(map[string]any)

	if len(responseBody) != 0 {
		err := json.Unmarshal(responseBody, &responseMap)
		if err != nil {
			log.Println("Fatal Error on Response Unmarshalling: ", err.Error());
			return nil, err
		}
	}

	return responseMap, nil
}

// Handle any of the Vault API Responses - ** Needs Work
func handleVaultAPIResponse(apiResponse map[string]any, messagePrefix string) (map[string]any, error) {
	if apiResponse["errors"] != nil {
		errorMessages := apiResponse["errors"].([]interface{})

		for _, value := range errorMessages {
			return nil, fmt.Errorf("%v Response Error: %v", messagePrefix, value)
		}

		return nil, nil

	} else if apiResponse["warnings"] != nil{
		warningMessages := apiResponse["warnings"].([]interface{})

		for _, value := range warningMessages {
			return nil, fmt.Errorf("%v Response Warning: %v", messagePrefix, value)
		}

		return nil, nil

	} else if apiResponse["data"] == nil {
		return nil, fmt.Errorf("%v Nil Reponse Error", messagePrefix)
	} else {
		return apiResponse, nil
	}
}
