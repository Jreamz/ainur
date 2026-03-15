## Ainur

## Project setup notes
1. `mkdir ainur`
2. `cd ainur`
3. `mkdir{cmd, internal, static}`
4. `go mod init ainur`
5. go get -u github.com/go-chi/chi/v5 (https://github.com/go-chi/chi)
6. Downloaded htmx.min.js.js (https://htmx.org/) into static dir
7. Downloaded pico css directly (https://picocss.com/docs/version-picker) into static dir

## Flow

Browser → POST → Chi handler → reads form data
↓
provision.go → POST → external API
↓
gets response back
↓
handler renders result HTML
↓
Browser ← HTML fragment back
