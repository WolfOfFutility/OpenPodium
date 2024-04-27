package main

// ** Dependency Documentation:
// User Auth: https://developer.hashicorp.com/vault/api-docs/auth/userpass

// ** File Breakdown:
// This is intended to create a secure store for the end-user to store secrets, Personal Access Tokens and SSO Tokens to allow
// easy integration into a host of tools while maintaining some sembelance of security. This includes C.R.U.D. functionality for
// the backend of the secure store.

import (
	"fmt"
	"log"
	"net/http"
)

// Enable User/Pass Auth at the HashiCorp Vault
func enableUserAuth() {
	// Create the headers and the body
	apiHeaders := http.Header {
		"X-Vault-Token": {"my-token"},
	}

	apiBody := map[string]any {
		"type": "userpass",
	}

	// Send the API to enable user auth
	response, err := makeAPIRequest("POST", "http://127.0.0.1:8200/v1/sys/auth/userpass", apiHeaders, apiBody)
	if err != nil {
		log.Println("Enable App Roles Error: ", err.Error())
	}
	
	// Handle the errors or warnings in the response
	if response["errors"] != nil {
		errorMessages := response["errors"].([]interface{})

		for _, value := range errorMessages {
			log.Println("Enable User Auth Error Message: ", value)
		}
	} else {
		log.Println("Enable User Auth Response: ", response)
	}
}

// Login to the HashiCorp Vault with a Username and Password
func loginVaultUser(login Login) (AuthUser, error) {
	if login.Username == "root" && login.Password == "root" { // Allows for root login (for now)
		return AuthUser{Username: login.Username, ClientToken: "my-token"}, nil
	} else {
		// Builds the API Request for the login
		apiBody := map[string]any {
			"password": login.Password,
		}
	
		// Sends the API Request, parses the response
		loginUserResponse, err := makeAPIRequest("POST", fmt.Sprintf("http://127.0.0.1:8200/v1/auth/userpass/login/%v", login.Username), nil, apiBody)
		if err != nil {
			log.Println("Login User Response Error: ", err.Error())
			return AuthUser{}, err
		}

		// Create the Auth User object so that the Client Token is passed.
		authUser := AuthUser{
			Username: login.Username,
			ClientToken: loginUserResponse["auth"].(map[string]interface{})["client_token"].(string),
		}
	
		return authUser, nil
	}
}

// Create a User within a HashiCorp Vault
func createVaultUser(loginToCreate Login, loginToAdminister Login) (error) {
	// Login to the Vault, get the client token
	auth, err := loginVaultUser(loginToAdminister)
	if err != nil {
		log.Println("Create User Error - Getting Client Token: ", err.Error())
		return err
	}

	// Set the User's Default Policy Settings
	defaultPolicySettings := PolicySettings {
		name: fmt.Sprintf("%v_user_default_policy", loginToCreate.Username),
		path: fmt.Sprintf("secret/data/%v/*", loginToCreate.Username),
		capabilities: "[\"read\", \"update\", \"list\", \"create\", \"delete\"]",
	}

	// Set API Headers and Body Data
	apiHeaders := http.Header {
		"X-Vault-Token": {auth.ClientToken},
		"Content-Type": {"application/json"},
	}

	apiBody := map[string]any {
		"password": loginToCreate.Password,
		"token_policies": defaultPolicySettings.name,
	}

	// Send the API Request - Create the User
	createUserResponse, err := makeAPIRequest("POST", fmt.Sprintf("http://127.0.0.1:8200/v1/auth/userpass/users/%v", loginToCreate.Username), apiHeaders, apiBody)
	if err != nil {
		log.Println("Create User Response Error: ", err.Error())
		return err
	}

	// Create the Vault Policy for the user
	// This will give them their own secret area to manage in the vault
	vaultPolicyError := createVaultPolicy(auth, defaultPolicySettings.name, defaultPolicySettings.path, defaultPolicySettings.capabilities)
	if vaultPolicyError != nil {
		log.Println("Policy Creation Error: ", vaultPolicyError.Error())
	}

	log.Println(createUserResponse)

	return nil
}

// Create a HashiCorp Vault Secret Policy
func createVaultPolicy(auth AuthUser, policyName string, pathName string, capabilities string) (error) {
	apiHeaders := http.Header {
		"X-Vault-Token": {auth.ClientToken},
		"Content-Type": {"application/json"},
	}

	apiBody := map[string]any {
		"policy": fmt.Sprintf("path \"%v\" {\"capabilities\" = %v}", pathName, capabilities),
	}

	createPolicyResponse, err := makeAPIRequest("POST", fmt.Sprintf("http://127.0.0.1:8200/v1/sys/policies/acl/%v", policyName), apiHeaders, apiBody)
	if err != nil {
		log.Println("New App Secret ID Error: ", err.Error())
		return err
	}

	log.Println("Create Policy Response: ", createPolicyResponse)

	return nil
}

// Create or Update a HashiCorp Vault Secret
func createVaultSecret(auth AuthUser, secretName string, secretValue map[string]any) (error) {
	//  Read the original secret value
	currentValue, errCode, err := readVaultSecret(auth, secretName);
	if err != nil && errCode != 100 {
		log.Println("Create Secret Read Error: ", err.Error())
		return err
	} 
	
	// Filter if there is no secret found
	if errCode == 100 {
		log.Printf("No Secret Found, defaulting to empty.")
	}

	// Add more secret values to be saved
	for name, value := range secretValue {
		currentValue[name] = value
	}

	// Set API Headers and Body Data
	apiHeaders := http.Header {
		"X-Vault-Token": {auth.ClientToken},
		"Content-Type": {"application/json"},
	}

	// Build the API Body
	apiBody := map[string]any {
		"data": currentValue,
	}

	// Send API Secret Creation Request
	createSecretResponse, err := makeAPIRequest("POST", fmt.Sprintf("http://127.0.0.1:8200/v1/secret/data/%v/%v", auth.Username, secretName), apiHeaders, apiBody)
	if err != nil {
		log.Println("New App Secret ID Error: ", err.Error())
		return err
	}

	// Handle Errors or Warnings from the API Response Body
	if createSecretResponse["errors"] != nil {
		errorMessages := createSecretResponse["errors"].([]interface{})

		for _, value := range errorMessages {
			log.Println("Create Secret Response Error: ", value)
		}

		return nil

	} else if createSecretResponse["warnings"] != nil{
		warningMessages := createSecretResponse["warnings"].([]interface{})

		for _, value := range warningMessages {
			log.Println("Create Secret Response Warning: ", value)
		}

		return nil

	} else {
		log.Println(createSecretResponse)
		return nil
	}
}

// Read a HashiCorp Vault Secret
func readVaultSecret(auth AuthUser, secretName string) (map[string]any, int32, error) {
	// Set API Headers and Body Data
	apiHeaders := http.Header {
		"X-Vault-Token": {auth.ClientToken},
		"Content-Type": {"application/json"},
	}

	// Send Read Secret API, handle errors
	readSecretResponse, err := makeAPIRequest("GET", fmt.Sprintf("http://127.0.0.1:8200/v1/secret/data/%v/%v", auth.Username, secretName), apiHeaders, nil)
	if err != nil {
		log.Println("Read Secret Error: ", err.Error())
		return nil, 400, err
	}

	if readSecretResponse["data"] == nil {
		return map[string]any{}, 100, fmt.Errorf("read secret data is nil")
	}

	// Parse API Response, send values back
	return readSecretResponse["data"].(map[string]interface{})["data"].(map[string]interface{}), 200, nil
}

// Delete a HashiCorp Vault Secret
func deleteVaultSecret(auth AuthUser, secretName string) (error) {
	// Set API Headers and Body Data
	apiHeaders := http.Header {
		"X-Vault-Token": {auth.ClientToken},
		"Content-Type": {"application/json"},
	}

	// Send the API Request to delete the specified secret
	deleteSecretResponse, err := makeAPIRequest("DELETE", fmt.Sprintf("http://127.0.0.1:8200/v1/secret/data/%v/%v", auth.Username, secretName), apiHeaders, nil)
	if err != nil {
		log.Println("Delete Secret Error: ", err.Error())
		return err
	}

	log.Println(deleteSecretResponse)
	return nil
}