package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	const songDir = "songs"
	const port = 8080

	http.Handle("/",AddHeaders(http.FileServer(http.Dir(songDir))))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func AddHeaders(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		h.ServeHTTP(w, r)
	}
}
