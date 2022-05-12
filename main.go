package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

func main() {
	// handle route
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintf(w, "<h1>Hello World!</h1>")
	})

	// print a log message
	log.Println("Starting server on port", webPort)

	// start the server
	http.ListenAndServe(fmt.Sprintf(":%s", webPort), nil)
}
