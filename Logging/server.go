package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("This is some text"))
}

func main(){
	addr := flag.String("addr",":8080","Network Address")
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "Error\t", log.Ldate|log.Ltime|log.Lshortfile)
	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("/",home)

	srv := &http.Server{
		Addr: *addr,
		ErrorLog: errorLog,
		Handler: mux,
	}
	infoLog.Print("Start server at",*addr)
	err := srv.ListenAndServe()

	errorLog.Fatal(err)
}