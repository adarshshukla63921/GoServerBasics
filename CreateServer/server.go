package main

import (
	"log"
	"net/http"
)

// creating our handler
func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this is html server"))
}

// creating our servemux and running our server
func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/",home)
	log.Print("Starting server at port 8080")
	err:= http.ListenAndServe("localhost:8080",mux)
	log.Fatal(err)
}