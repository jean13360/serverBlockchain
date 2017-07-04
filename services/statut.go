package Services

import (
	"fmt"
	"html"
	"net/http"
)

//Option return header
func Option(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
}

//Status return if server is alive
func Status(w http.ResponseWriter, r *http.Request) {
	Option(w, r)
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}
