package main

// run and visit
// http://localhost:8800/api/v201910/?apikey=APIKEY&product=PRODUCT&uuid=UUID

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	apikey := r.FormValue("apikey")
	product := r.FormValue("product")
	uuid := r.FormValue("uuid")
	version := r.FormValue("version")
	fmt.Printf("apikey %s, product %s, version %s, uuid %s, path %s\n",
		apikey, product, version, uuid, path)
}

func main() {
	http.HandleFunc("/api/v201910/", apiHandler)

	port := os.Getenv("INGA_PORT")
	if port == "" {
		port = "8800"
	}
	port = ":" + port
	fmt.Fprintln(os.Stderr, "Listening on", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
