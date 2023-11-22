package requests

import (
	"fmt"
	"net/http"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello, World!<h1>")
}
