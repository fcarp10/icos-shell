# icos-shell - CLI

ICOS Shell CLI

## Development 

### Client API
```
openapi-generator-cli generate -g go -i openapi.yaml -o client/pkg/openapi --additional-properties=packageName=openapi,isGoSubmodule=true

rm client/pkg/openapi/go.mod client/pkg/openapi/go.sum
```

### Docs
```
go run main.go create docs --path docs
```

### Building with Docker
For convenience and reproducibility, a build script based on Docker has been added, as well as a Dockerfile to accompany it. To use it, make sure you have Docker installed and simply execute the script from within this folder via ```./build.sh```. The compiled binary will then be placed in this folder (old binaries will be overridden).

Documentation [here](./docs/icos-shell.md)
