package handlers

import (
	"net/http"
)

func AiWebsite(response http.ResponseWriter, request *http.Request) {
	http.ServeFile(response, request, "./myAIwebsite/ai.html")
}
