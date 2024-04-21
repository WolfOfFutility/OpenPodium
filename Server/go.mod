module github.com/WolfOfFutility/OpenPodium/Server

go 1.22.2

require (
	github.com/google/uuid v1.6.0
	github.com/hashicorp/vault-client-go v0.4.3
)

require (
	// github.com/WolfOfFutility/OpenPodium/Server/handlers v0.1.2
	// github.com/WolfOfFutility/OpenPodium/Server/testers v0.1.0
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-retryablehttp v0.7.5 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/go-secure-stdlib/strutil v0.1.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/ryanuber/go-glob v1.0.0 // indirect
	golang.org/x/sys v0.19.0 // indirect
	golang.org/x/time v0.5.0 // indirect
)

replace github.com/WolfOfFutility/OpenPodium/Server/handlers => ./handlers
