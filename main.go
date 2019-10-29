package main

// run with
// ./inga
// or specify a different port (the default is 8800 for development)
// INGA_PORT=8080 ./inga
// and visit
// http://localhost:8800/api/v201910/?apikey=APIKEY&product=PRODUCT&uuid=UUID&eventID=EVENTID&dateTime=DATETIME

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	apikey := r.FormValue("apikey")
	product := r.FormValue("product")
	uuid := r.FormValue("uuid")
	version := r.FormValue("version")
	eventID := r.FormValue("eventID")
	dateTime := r.FormValue("dateTime")
	fmt.Printf("apikey %s, product %s, version %s, uuid %s, eventID %s, dateTime %s, path %s\n",
		apikey, product, version, uuid, eventID, dateTime, path)
}

func main() {
	http.HandleFunc("/api/v201910/", apiHandler)

	// Detect systemd.
	// https://www.darkcoding.net/software/systemd-socket-activation-in-go/
	if os.Getenv("LISTEN_PID") == strconv.Itoa(os.Getpid()) {
		// systemd
		fd := os.NewFile(3, "from systemd")
		l, err := net.FileListener(fd)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(os.Stderr, "Listening via systemd")
		http.Serve(l, nil)
	} else {
		// not systemd (manual)
		port := os.Getenv("INGA_PORT")
		if port == "" {
			port = "8800"
		}
		port = ":" + port
		fmt.Fprintln(os.Stderr, "Listening on", port)
		log.Fatal(http.ListenAndServe(port, nil))
	}
}
