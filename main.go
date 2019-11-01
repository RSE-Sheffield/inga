package main

// run with
// ./inga
// or specify a different port (the default is 8800 for development)
// INGA_PORT=8080 ./inga
// and visit
// http://localhost:8800/api/v201910/?apikey=APIKEY&product=PRODUCT&uuid=UUID&eventID=EVENTID&dateTime=DATETIME

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

var logFile io.Writer

func apiHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	apikey := r.FormValue("apikey")
	product := r.FormValue("product")
	uuid := r.FormValue("uuid")
	version := r.FormValue("version")
	eventID := r.FormValue("eventid")
	dateTime := r.FormValue("datetime")
	fmt.Fprintf(logFile, "%s, %s, %s, %s, %s, %s, %s\n",
		apikey, product, version, uuid, eventID, dateTime, path)
}

func pinga() {
	// pings the inga server when this instance of inga is started
	target := "http://inga.shef.ac.uk:80/api/v201910/"
	fmt.Fprintf(os.Stderr, "Pinging inga @ %s ...\n", target)

	form := url.Values{}
	form.Add("apikey", "ak-sample")
	form.Add("product", "inga")
	form.Add("uuid", "inga")
	form.Add("version", "v201910")
	form.Add("eventid", "inga")
	t := time.Now()
	form.Add("datetime", t.Format("20060102150405"))
	resp, err := http.PostForm(target, form)
	defer resp.Body.Close()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	// creates a new log file with a timestamped name
	t := time.Now()
	fname := "inga_" + t.Format("20060102150405") + ".log"
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0444)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	logFile = io.Writer(f)
	fmt.Fprintf(logFile, "apikey, product, version, uuid, eventid, datetime, path\n") //print csv header

	http.HandleFunc("/api/v201910/", apiHandler)

	port := os.Getenv("INGA_PORT")
	if port == "" {
		port = "8800"
	}
	port = ":" + port
	fmt.Fprintln(os.Stderr, "Listening on", port)
	go pinga()
	log.Fatal(http.ListenAndServe(port, nil))
}
