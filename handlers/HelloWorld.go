package handlers

import (
	"fmt"
	"net/http"
)

func HelloWorld(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	fmt.Fprint(response, "<h1>Hello, World!<h1>")
}
