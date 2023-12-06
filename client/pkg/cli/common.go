package cli

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/hokaccha/go-prettyjson"
	"github.com/spf13/viper"
)

func printError(resp *http.Response, err error) bool {
	if err != nil {
		if resp != nil {
			fmt.Fprintln(os.Stderr, "Error:", resp.Body)
		}
		fmt.Fprintln(os.Stderr, "Error:", err)
		return true
	}
	fmt.Fprintln(os.Stdout, resp.StatusCode)
	return false
}

func printResponseSimple(resp *http.Response, err error) {
	if !printError(resp, err) {
		fmt.Fprintln(os.Stderr, resp.Body)
	}
}

func printPrettyJSON(object map[string]interface{}, resp *http.Response, err error) {
	if !printError(resp, err) {
		if resp.StatusCode == 200 || resp.StatusCode == 201 {
			prettyJSON, err := prettyjson.Marshal(object)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error marshaling JSON:", err)
				return
			} else {
				fmt.Fprintln(os.Stderr, string(prettyJSON))
			}
		} else {
			fmt.Fprintln(os.Stderr, resp.Body)
		}
	}
}

func printArrayPrettyJSON(object []map[string]interface{}, resp *http.Response, err error) {
	for _, value := range object {
		printPrettyJSON(value, resp, err)
	}
}

func CleanToken() {
	token := viper.GetString("auth_token")
	token = strings.ReplaceAll(token, "\n", "")
	fmt.Fprintln(os.Stderr, "Token found:", "..."+token[len(token)-50:len(token)-1])
	viper.Set("auth_token", token)
}
