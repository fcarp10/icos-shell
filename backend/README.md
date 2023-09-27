# shell-backend

ICOS Shell backend based on the OpenAPI 3.0 specification.

## Development 

### Backend API
```
openapi-generator-cli generate -g go-server -i openapi.yaml -o backend/ --additional-properties=packageName=shellbackend
```

### Docs
```
openapi-generator-cli generate -g markdown -i openapi.yaml -o backend/docs/
```

Documentation [here](./docs/README.md)