package main

import (
	"log"
	"net/http"
)

// creating different hanlders for diffrent routes
func home(w http.ResponseWriter, r *http.Request){
	if r.URL.Path != "/" {
		http.NotFound(w,r)
		return
	}
	w.Write([]byte("this is for home"))
}
func createUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this sections is for user creation"))
}

func editUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this seciton is for altering existing users"))
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("this seciton is for deleting existing users"))
}


func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/",home)
	mux.HandleFunc("/user/create",createUser)
	mux.HandleFunc("/user/edit",editUser)
	mux.HandleFunc("/user/delete",deleteUser)

	log.Print("Starting server at localhost:8080")
	err := http.ListenAndServe("localhost:8080",mux)
	log.Fatal(err)
}