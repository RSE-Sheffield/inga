package main

import (
	"fmt"
	"log"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path
	fmt.Printf("%s\n", title)
}

func main() {
	http.HandleFunc("/api/v201910/", apiHandler)
	log.Fatal(http.ListenAndServe(":8800", nil))
}
