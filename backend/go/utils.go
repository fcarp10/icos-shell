package shellbackend

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func prepareToken(ctx context.Context, apiKey string, req *http.Request) *http.Request {
	_, authToken, _ := receiveAndValidateAccessToken(ctx, apiKey)
	authToken = "Bearer " + authToken
	req.Header.Add("Authorization", authToken)
	return req
}

func errorConnect(resp *http.Response, err error) (ImplResponse, error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	if resp != nil {
		return Response(500, resp.Body), nil
	}
	return Response(500, "Error while connecting to the component"), nil
}

func unexpectedCode(statusCode int) (ImplResponse, error) {
	return Response(statusCode, "Unexpected status code received"), nil
}

func readResponse(response *http.Response) []byte {
	body, err := io.ReadAll(response.Body)
	response.Body.Close()
	if err != nil {
		log.Println("Error reading response body:", err)
		return nil
	} else {
		// log.Println(response.StatusCode, string(body))
		return body
	}
}

func unmarshalArrayResponse(response *http.Response) []map[string]interface{} {
	var body = readResponse(response)
	var data []map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return nil
	}
	return data
}

func unmarshalResponse(response *http.Response) map[string]interface{} {
	var body = readResponse(response)
	var data map[string]interface{}
	err := json.Unmarshal(body, &data)
	if err != nil {
		log.Println("Error unmarshalling response:", err)
		return nil
	}
	return data
}
