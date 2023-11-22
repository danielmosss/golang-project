package requests

import (
	"net/http"
)

func HtmlHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./myAIwebsite/ai.html")
}
