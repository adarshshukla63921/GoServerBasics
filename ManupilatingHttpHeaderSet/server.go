package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	// header must be set before write.
	w.Header().Set("Cache-Control","public, max-age = 12456")
	w.Header().Add("Cache-Control","new cache header")
	w.Header().Add("Cache-Control","last cache header")
	w.Header()["Date"] = nil
	w.Header()["Content-Length"] = nil
	
	// only write after you are done with headers, otherwise headers wont actually reach the client.
	w.Write([]byte("Here we learn about more functions in the w.Header()"))
	log.Print(w.Header().Values("Cache-Control"))
	log.Print("first value assigned to header :",w.Header().Get("Cache-Control"))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	err := http.ListenAndServe("localhost:8080",mux)
	log.Fatal(err)
}