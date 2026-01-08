package main

import (
	"log"
	"net/http"
)

func postMethodHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		w.Header().Set("Allow",http.MethodPost)
		http.Error(w,"Only Post is allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("content-type","application/json")
	w.Write([]byte(`{
		"name" : "adarsh",
		"branch" : "CSAI",
		"course" : "Btech",
		"interest" : "Android Development"
	}`))
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/createPost",postMethodHandler)
	log.Print("Starting the server at 8080")
	err := http.ListenAndServe("localhost:8080",mux)

	log.Fatal(err)
}