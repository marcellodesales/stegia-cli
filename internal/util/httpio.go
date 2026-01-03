package util

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

func BasicAuthHeader(userPass string) string {
	return "Basic " + base64.StdEncoding.EncodeToString([]byte(userPass))
}

func PrintHTTPRequest(req *http.Request, body []byte) {
	fmt.Println("\n=== HTTP REQUEST ===")
	fmt.Printf("%s %s\n", req.Method, req.URL.String())
	for k, v := range req.Header {
		fmt.Printf("%s: %s\n", k, strings.Join(v, ", "))
	}
	fmt.Println()
	if len(body) > 0 {
		fmt.Println(string(body))
	}
}

func PrintHTTPResponse(statusCode int, headers map[string]string, body []byte) {
	fmt.Println("\n=== HTTP RESPONSE ===")
	fmt.Printf("HTTP/1.1 %d\n", statusCode)
	for k, v := range headers {
		fmt.Printf("%s: %s\n", k, v)
	}
	fmt.Println()
	if len(body) > 0 {
		fmt.Println(string(body))
	}
}
