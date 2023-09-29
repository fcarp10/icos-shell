# ICOS Shell

ICOS Shell is a component part of the [ICOS Project](https://cordis.europa.eu/project/id/101070177)

- Backend: server running on the ICOS controller
- Client: CLI tool for interfacing the Shell backend

## Build and run

### Backend (with Docker)
```
cd backend
docker build -t "shell-backend" .
docker run --volume ./config.yml:/app/config.yml -p 8080:8080 shell-backend:latest
```
More info [here](./backend/README.md)

### Client - CLI
```
cd client
go build -o icos-shell
./icos-shell -h
```
More info [here](./client/README.md)

