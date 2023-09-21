# ICOS Shell

ICOS Shell is a component part of the [ICOS Project](https://cordis.europa.eu/project/id/101070177)

- Backend: server running on the ICOS controller
- Client: CLI tool for interfacing the Shell backend

## Run

### Backend
```
docker run fcarp10/shell-backend
```

### CLI
```
icos-cli -h
NAME:
   icos-cli - CLI

USAGE:
   icos-cli [global options] command [command options] [arguments...]

VERSION:
   v0.1

COMMANDS:
   controller, c  options for controllers
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config value, -c value  config file (default: config.yaml) [$CONFIG_FILE]
   --server value            URL of the shell-backend (default: "localhost:8080")
   --username value          username (default: "admin")
   --password value          password
   --help, -h                show help
   --version, -v             print the version
```


## Build

### Backend
```
cd backend
docker build .
```

### CLI
```
cd client
go build -o icos-cli
```

### Generate OpenAPI

#### Server
```
openapi-generator-cli generate -g go-server -i openapi.yaml -o backend/ --additional-properties=packageName=shellbackend
```

#### Client
```
openapi-generator-cli generate -g go -i openapi.yaml -o client/openapi --additional-properties=packageName=shellclient,isGoSubmodule=true

rm client/openapi/go.mod client/openapi/go.sum
```

#### Docs
```
openapi-generator-cli generate -g markdown -i openapi.yaml -o docs/
```


