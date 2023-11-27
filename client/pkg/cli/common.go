package cli

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func printResponseSimple(resp *http.Response, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		fmt.Fprintln(os.Stderr, resp.Body)
		fmt.Fprintln(os.Stdout, resp.StatusCode)
	}
}

func printPrettyJSON(object map[string]interface{}, resp *http.Response, err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		if resp != nil {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	} else {
		if resp.StatusCode == 200 {
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
