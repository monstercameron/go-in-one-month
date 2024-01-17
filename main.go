package main

import (
	"net/http"
)

func main() {
	// server templ file to / route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/index.html")
	});
}