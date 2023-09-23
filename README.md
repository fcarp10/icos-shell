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
./icos-shell -h
icos-shell - a CLI to interface the ICOS Shell
   
The icos-shell can be used to modify or inspect resources in the ICOS controller from the terminal

Usage:
  icos-shell [command]

Available Commands:
  add         Add a resource
  completion  Generate the autocompletion script for the specified shell
  get         Get a resource
  help        Help about any command

Flags:
      --config string     config file (default "$XDG_CONFIG_HOME/icos-shell/config.yaml")
  -h, --help              help for icos-shell
  -p, --password string   password parameter
  -s, --server string     server URL (default "localhost:8080")
  -u, --username string   username parameter (default "admin")

Use "icos-shell [command] --help" for more information about a command.
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
go build -o icos-shell
```

### Generate OpenAPI

#### Server
```
openapi-generator-cli generate -g go-server -i openapi.yaml -o backend/ --additional-properties=packageName=shellbackend
```

#### Client
```
openapi-generator-cli generate -g go -i openapi.yaml -o client/pkg/openapi --additional-properties=packageName=openapi,isGoSubmodule=true

rm client/openapi/go.mod client/openapi/go.sum
```

#### Docs
```
openapi-generator-cli generate -g markdown -i openapi.yaml -o docs/
```


