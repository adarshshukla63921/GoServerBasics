package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	// here we are trying to get the value of id from our request's URL
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		// in-case we cannot get an id, or get an invalid id we show this
		http.NotFound(w,r)
		return
	}
	// this is the 'page' we see when our id is valid
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
	log.Print(id)
}

func main() {
	mux := http.NewServeMux()

	// serve matches only the url path, not the query string, it is parsed after routing.
	mux.HandleFunc("/view",home)

	err := http.ListenAndServe("localhost:8080",mux)

	log.Fatal(err)
}