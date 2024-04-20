package main

// ** Dependency Documentation:
// Vault: https://github.com/hashicorp/vault-client-go

// ** File Breakdown:
// This is intended to create a secure store for the end-user to store secrets, Personal Access Tokens and SSO Tokens to allow 
// easy integration into a host of tools while maintaining some sembelance of security. This includes C.R.U.D. functionality for
// the backend of the secure store.

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/vault-client-go"
	"github.com/hashicorp/vault-client-go/schema"
)

// Handles the C and U functions in C.R.U.D. for the user's secret vault
func writeSecretToUserVault(userName string, userSecret map[string]any) {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
	}

	// check the users current secret store
	userSecretList, err := readSecretsFromUserVault(userName)
	if err != nil {
		log.Fatal(err)
	}

	// for each of the secrets sent through, save them to the user secret list variable
	for name, value := range userSecret {
		userSecretList[name] = value
	}

	// declare the secret list and make the changes to the user's
	secretList := make(map[string]map[string]any)
	secretList[userName] = userSecretList

	// write secret list to vault
	_, err = client.Secrets.KvV2Write(ctx, "user-secrets", schema.KvV2WriteRequest{Data: map[string]any {"secretList" : secretList}},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log.Println("secret written successfully")
}

// Handles the R functions in the C.R.U.D. for the user's secret vault
func readSecretsFromUserVault(userName string) (map[string]any, error) {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
		return nil, err
	}

	// read the secret
	s, err := client.Secrets.KvV2Read(ctx, "user-secrets", vault.WithMountPath("secret"))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// parse the secret store
	secretList := s.Data.Data["secretList"].(map[string]any)

	// If there are no existing secrets, return an empty secret store for the user
	// If there are existing secrets for the user, return them
	if secretList[userName] == nil {
		log.Println("no secrets could be found for user.")
		return make(map[string]any), nil
	} else {
		// log.Println("secrets retrieved:", secretList[userName])
		return secretList[userName].(map[string]any), nil
	}
}

// Handles the D functions in the C.R.U.D. functions in the user's secret vault
func removeSecretFromUserVault(userName string, secretName string) (string, error) {
	ctx := context.Background()

	// prepare a client with the given base address
	client, err := vault.New(
		vault.WithAddress("http://127.0.0.1:8200"),
		vault.WithRequestTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatal(err)
	}

	// authenticate with a root token (insecure)
	if err := client.SetToken("my-token"); err != nil {
		log.Fatal(err)
	}

	// check the users current secret store
	userSecretList, err := readSecretsFromUserVault(userName)
	if err != nil {
		log.Fatal(err)
	}

	// delete the specified index from the list
	delete(userSecretList, secretName)

	// declare the secret list and make the changes to the user's
	secretList := make(map[string]map[string]any)
	secretList[userName] = userSecretList

	// write secret list to vault
	_, err = client.Secrets.KvV2Write(ctx, "user-secrets", schema.KvV2WriteRequest{Data: map[string]any {"secretList" : secretList}},
		vault.WithMountPath("secret"),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Successfully removed secret: ", secretName)
	return "successfully removed secret", nil
}