package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func printError(resp *http.Response, err error) bool {
	if err != nil {
		if resp != nil {
			fmt.Fprintln(os.Stderr, "Error:", resp.Body)
		}
		fmt.Fprintln(os.Stderr, "Error:", err)
		return true
	}
	return false
}

func printResponseSimple(resp *http.Response, err error) {
	if !printError(resp, err) {
		fmt.Fprintln(os.Stderr, resp.Body)
		fmt.Fprintln(os.Stdout, resp.StatusCode)
	}
}

func printPrettyJSON(object map[string]interface{}, resp *http.Response, err error) {
	if !printError(resp, err) {
		if resp.StatusCode == 200 || resp.StatusCode == 201 {
			prettyJSON, err := json.MarshalIndent(object, "", "  ")
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
				return
			} else {
				fmt.Fprintln(os.Stdout, string(prettyJSON))
			}
		} else {
			fmt.Fprintln(os.Stderr, resp.Body)
			fmt.Fprintln(os.Stdout, resp.StatusCode)
		}
	}
}

func printArrayPrettyJSON(object []map[string]interface{}, resp *http.Response, err error) {
	for _, value := range object {
		printPrettyJSON(value, resp, err)
	}
}
