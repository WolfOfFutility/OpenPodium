Set-Location ./Server

vault server -dev -dev-root-token-id="my-token"
go run .