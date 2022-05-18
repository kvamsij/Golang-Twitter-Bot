package main

import (
	"bet365/routeHandlers"
	"net/http"
)

func main() {
	http.HandleFunc("/twitterbot", routeHandlers.TwitterBot)
	http.ListenAndServe("localhost:8080", nil)
}
