package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	httpPort := flag.String("p", "3000", "Port for localhost")
	directory := flag.String("d", ".", "Folder containing static files to serve")
	flag.Parse()

	// Validate path
	absolutePath, err := filepath.Abs(*directory)
	if err != nil {
		log.Fatal(err)
	}

	// Using http.DefaultServeMux is not best practice, as it's globally scoped
	// We will use a local mux instead (and DefaultServeMux is simply an instance of NewServeMux())
	mux := http.NewServeMux()

	mux.Handle("/", http.FileServer(http.Dir(absolutePath)))

	log.Printf("Serving %s as http://localhost:%s\n", absolutePath, *httpPort)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *httpPort), logEveryRequest(mux)))
}

func logEveryRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		log.Printf("%s %s %s\n", request.RemoteAddr, request.Method, request.URL)
		handler.ServeHTTP(writer, request)
	})
}
