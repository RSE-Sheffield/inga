package main

// run and visit
// http://localhost:8800/api/v201910/?apikey=APIKEY&product=PRODUCT&uuid=UUID

import (
	"fmt"
	"log"
	"net/http"
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
	log.Fatal(http.ListenAndServe(":8800", nil))
}
