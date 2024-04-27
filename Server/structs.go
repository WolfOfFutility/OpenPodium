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