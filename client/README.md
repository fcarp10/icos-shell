# icos-shell - CLI

ICOS Shell CLI

## Development 

### Client API
```
openapi-generator-cli generate -g go -i openapi.yaml -o client/pkg/openapi --additional-properties=packageName=openapi,isGoSubmodule=true

rm client/openapi/go.mod client/openapi/go.sum
```

### Docs
```
go run main.go create docs --path docs
```

Documentation [here](./docs/icos-shell.md)