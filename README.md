# ICOS Shell

ICOS Shell is a component part of the [ICOS Project](https://cordis.europa.eu/project/id/101070177)

## Generate OpenAPI

Generate the API from the backend with:
```
openapi-generator-cli generate -g go-server -i openapi.yaml -o backend/ --additional-properties=packageName=shellbackend
```

Generate the API from the client with:
```
openapi-generator-cli generate -g go -i openapi.yaml -o client/openapi --additional-properties=packageName=shellclient,isGoSubmodule=true

rm client/openapi/go.mod client/openapi/go.sum
```

Generate the docs for the APIs with:
```
openapi-generator-cli generate -g markdown -i openapi.yaml -o docs/
```

## Testing the lighthouse

Retrieve the list of controllers from the lighthouse with:
```
curl -XGET http://lighthouse.icos-project.eu:8080/api/v3/controller/
```

Add a new controller to the lighthouse with:
```
curl -XPOST 'http://lighthouse.icos-project.eu:8080/api/v3/controller/?username=admin&password=Iki946D56!!J@gSHpuonoUyH1uB*^' -d '{"name": "controller_1", "address": "192.168.100.2"}'
```

## Run

Run the backend with:

```
docker run fcarp10/shell-backend
```

## Debug

### Backend

Adjust the relative paths from `backend/main.go`:

```
	shellbackend "github.com/GIT_USER_ID/GIT_REPO_ID/go" // full path for debugging
	// shellbackend "./go" // comment out for debugging
```

And run the server with:
```
cd backend
go run main.go
```

### Client

Assuming the server is running in localhost, run the client with:
```
cd client
go run main.go
2023/09/19 14:10:04 OK!
```


## License
TBD
