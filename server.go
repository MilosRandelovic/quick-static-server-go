package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	httpPort := flag.String("p", "3000", "Port for localhost")
	directory := flag.String("d", ".", "Folder containing static files to serve")
	flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))

	log.Printf("Serving %s on HTTP port: %s\n", *directory, *httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *httpPort), logRequest(http.DefaultServeMux)))
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
