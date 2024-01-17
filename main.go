package main

import (
	"github.com/monstercameron/go-in-one-month/views/pages/index"
	"net/http"
)

func main() {
	// server templ file to / route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		component := views.IndexPage("Cam Cameron")
		// serve text/html
		w.Header().Set("Content-Type", "text/html")
		// render the component to the response writer
		component.Render(r.Context(), w)
	})

	http.ListenAndServe(":3000", nil)
}
