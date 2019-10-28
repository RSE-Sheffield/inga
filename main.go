package main

// run and visit
// http://localhost:8800/api/v201910/?api=APIKEY&product=PRODUCT&uuid=UUID

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
	fmt.Printf("apikey %s, product %s, uuid %s, path %s\n",
		apikey, product, uuid, path)
}

func main() {
	http.HandleFunc("/api/v201910/", apiHandler)
	log.Fatal(http.ListenAndServe(":8800", nil))
}
