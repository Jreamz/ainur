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


## Run Ainur server
`go run cmd/main.go`

## Run authentik-dev server
You will need docker and docker-compose installed on your local system

Just run the following:
```bash
sudo docker-compose pull
```
You should see the following:
```bash
m450n@ubuntu:~/Code/ainur/authentik-dev$ sudo docker-compose pull
Pulling postgresql ... done
Pulling server     ... downloading (81.2%)
Pulling worker     ... downloading (81.2%)
```
Then run the following to stand up authentik containers
```bash
sudo docker-compose up -d
```
Confirm authentik-dev containers are running:
```bash
m450n@ubuntu:~/Code/ainur/authentik-dev$ sudo docker-compose ps
           Name                         Command                 State                           Ports                    
-------------------------------------------------------------------------------------------------------------------------
authentik-dev_postgresql_1   docker-entrypoint.sh postgres   Up (healthy)   5432/tcp                                     
authentik-dev_server_1       dumb-init -- ak server          Up (healthy)   0.0.0.0:9000->9000/tcp,:::9000->9000/tcp,    
                                                                            0.0.0.0:9443->9443/tcp,:::9443->9443/tcp     
authentik-dev_worker_1       dumb-init -- ak worker          Up (healthy)   
```
## References
- Authentik API: https://pkg.go.dev/goauthentik.io/api/v3@v3.2026020.16#section-readme
- Authentik Dev docs: https://docs.goauthentik.io/developer-docs/
- Authentik Dev Docker docs: https://docs.goauthentik.io/install-config/install/docker-compose/