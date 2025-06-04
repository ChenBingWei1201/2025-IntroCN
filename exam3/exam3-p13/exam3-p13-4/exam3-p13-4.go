package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hola~")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "404 Not Found")
}

func main() {
	fmt.Println("Launching server...")

	hh := http.HandlerFunc(helloHandler)
	http.Handle("/hola", hh)

	fs := http.FileServer(http.Dir("."))
	http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Try to serve the file
		_, err := http.Dir(".").Open(r.URL.Path[1:])
		if err != nil {
			notFoundHandler(w, r)
			return
		}
		fs.ServeHTTP(w, r)
	}))
	http.ListenAndServeTLS("140.112.41.208:30202", "server.cer",
		"server.key", nil)
}
