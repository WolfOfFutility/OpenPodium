package main

type Login struct {
	Username string
	Password string
}

type AuthUser struct {
	Username string
	ClientToken string
}

type User struct {
	userName string
	SSOToken string
}

type PolicySettings struct {
	name string
	path string
	capabilities string
}

// type VaultAPIResponse struct {
// 	data map[string]any
// 	errors map[string]any
// 	warnings map[string]any
// }