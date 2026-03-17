## Ainur

## Project setup notes
1. `mkdir ainur`
2. `cd ainur`
3. `mkdir{cmd, internal, static}`
4. `go mod init ainur`
5. `go get -u github.com/go-chi/chi/v5` (https://github.com/go-chi/chi)
6. `Downloaded htmx.min.js.js` (https://htmx.org/) into static dir
7. `Downloaded pico css directly` (https://picocss.com/docs/version-picker) into static dir
8. `Download pico colors css directly` (https://picocss.com/docs/color-picker) into static dir
9. `go mod tidy`


## Run server
`go run cmd/main.go`


## References
- Authentik API: https://pkg.go.dev/goauthentik.io/api/v3@v3.2026020.16#section-readme
- Authentik Dev docs: https://docs.goauthentik.io/developer-docs/