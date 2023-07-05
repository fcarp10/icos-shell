# ICOS Shell

ICOS Shell is a component part of the [ICOS Project](https://cordis.europa.eu/project/id/101070177)

## Usage

To generate the backend:
```
openapi-generator-cli generate -g go-server -i openapi.yaml -o backend/ --additional-properties=packageName=shellbackend
```

To generate the client:
```
openapi-generator-cli generate -g go -i openapi.yaml -o client/openapi --additional-properties=packageName=shellclient,isGoSubmodule=true

rm client/openapi/go.mod client/openapi/go.sum
```

To generate the docs:
```
openapi-generator-cli generate -g markdown -i openapi.yaml -o docs/
```

## License
TBD
