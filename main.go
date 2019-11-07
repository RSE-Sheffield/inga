package main

// run with
// ./inga
// or specify a different port (the default is 8800 for development)
// INGA_PORT=8080 ./inga
// and visit
// http://inga.shef.ac.uk/api/v201910/?apikey=APIKEY&product=PRODUCT&uuid=UUID&eventID=EVENTID&dateTime=DATETIME

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var DefaultPort string = "8800"

var logFile io.Writer

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fields := []string{
		r.FormValue("apikey"),
		r.FormValue("product"),
		r.FormValue("uuid"),
		r.FormValue("version"),
		r.FormValue("eventid"),
		r.FormValue("datetime"),
	}
	var args []interface{}
	for i := range fields {
		args = append(args, strings.Map(remove_bad, fields[i]))
	}
	fmt.Fprintf(logFile, "%s,%s,%s,%s,%s,%s\n", args...)
}

func remove_bad(r rune) rune {
	// Remove runes from user input that are considered "bad".
	// All control characters,
	// and the CSV metacharacters «"» and «,».
	switch {
	case r < ' ', r == '"', r == ',':
		return -1
	}
	return r
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
	fmt.Fprintf(logFile, "apikey,product,version,uuid,eventid,datetime\n") //print csv header

	http.HandleFunc("/api/v201910/", apiHandler)

	port := os.Getenv("INGA_PORT")
	if port == "" {
		port = DefaultPort
	}
	port = ":" + port
	fmt.Fprintln(os.Stderr, "Listening on", port)
	go pinga()
	log.Fatal(http.ListenAndServe(port, nil))
}
