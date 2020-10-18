package routes

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello")
}

func routeHandler() {
	http.HandleFunc("/", test)
}