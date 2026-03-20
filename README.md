# Ainur

HARP - High Account Resource Provisioner

A lightweight internal user provisioning tool built with Go, Chi, HTMX, and Pico CSS. Ainur provides a web interface for creating and managing users in Authentik, which acts as the SSO identity provider for downstream services.

## Architecture
```
Browser (HTMX + Pico CSS)
    │
    ├── GET  /              → RootHandler       → renders home.html
    ├── GET  /search        → SearchHandler     → renders search.html
    ├── POST /provision     → ProvisionHandler  → client.CreateUserRequest() → Authentik API
    ├── GET  /search-results→ SearchResultsHandler → client.ListUsers()     → Authentik API
    ├── POST /users/{id}/reset-password → ResetPasswordHandler              → Authentik API
    └── DELETE /users/{id}  → DeleteUserHandler                             → Authentik API
```

### Flow
```
main.go
  ├── Sets up Chi router and middleware
  ├── Parses all templates and fragments
  ├── Creates Authentik API client
  ├── Registers routes with handlers (closures capture templates + client)
  └── Starts HTTP server

handlers.go
  ├── Parses form data or URL params from the request
  ├── Validates input
  ├── Calls client methods (client.go) to interact with Authentik
  └── Renders the appropriate HTMX fragment (success, failure, results)

client.go
  ├── Wraps the Authentik Go SDK
  ├── Configures API client with endpoint and bearer token
  └── Exposes methods: CreateUserRequest(), ListUsers(), etc.
```

### HTMX Pattern

1. User interacts with a form or button
2. HTMX sends a request to a Chi route(s) (no full page reload)
3. Handler ingests the request and calls the Authentik API
4. Handler renders an HTML fragment
5. HTMX swaps the fragment into the target div on the page (inner or outer)

## Project Structure
```
ainur/
├── cmd/
│   └── main.go          # entry point: router, middleware, server
├── internal/
│   ├── handlers.go      # HTTP handlers
│   └── client.go        # Authentik API client
├── static/
│   ├── css/             # Pico CSS, custom overrides
│   ├── templates/       # full page templates
│   ├── fragments/       # HTMX response fragments
│   ├── htmx.min.js
│   └── favicon.ico
├── authentik-dev/       # local Authentik docker-compose
├── go.mod
└── README.md
```

## Run Ainur Server
```bash
go run cmd/main.go
```

Server starts on `http://localhost:3000`

## Run Authentik Dev Server

Requires Docker and Docker Compose.

Pull the images:
```bash
sudo docker-compose pull
```

You should see something like:
```
Pulling postgresql ... done
Pulling server     ... downloading (81.2%)
Pulling worker     ... downloading (81.2%)
```

Start the containers:
```bash
sudo docker-compose up -d
```

Confirm containers are running:
```bash
sudo docker-compose ps
```
```
           Name                         Command                 State                           Ports
-------------------------------------------------------------------------------------------------------------------------
authentik-dev_postgresql_1   docker-entrypoint.sh postgres   Up (healthy)   5432/tcp
authentik-dev_server_1       dumb-init -- ak server          Up (healthy)   0.0.0.0:9000->9000/tcp,:::9000->9000/tcp,
                                                                            0.0.0.0:9443->9443/tcp,:::9443->9443/tcp
authentik-dev_worker_1       dumb-init -- ak worker          Up (healthy)
```

Navigate to `http://localhost:9000/if/flow/initial-setup/` to create the admin account. Then create an API token under Directory -> Tokens & App Passwords (choose API Token).

## References

- [Authentik Go SDK](https://pkg.go.dev/goauthentik.io/api/v3)
- [Authentik Developer Docs](https://docs.goauthentik.io/developer-docs/)
- [Authentik Docker Installation](https://docs.goauthentik.io/install-config/install/docker-compose/)
- [Chi Router](https://github.com/go-chi/chi)
- [HTMX](https://htmx.org/)
- [Pico CSS](https://picocss.com/)
