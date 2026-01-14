package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func (app *application) home(w http.ResponseWriter, r *http.Request){
	app.logInfo.Print("Home was visited")
	w.Write([]byte("this is home page"))
}

func (app *application) aboutUs(w http.ResponseWriter, r *http.Request){
	app.logInfo.Print("About Us was visited.")
	w.Write([]byte("this is About-us page"))
}

func (app *application) view(w http.ResponseWriter, r *http.Request){
	app.logInfo.Print("View was visited.")
	w.Write([]byte("this is View-page page"))
}

type application struct{
	logError  *log.Logger
	logInfo    *log.Logger
}

func main(){
	addr := flag.String("addr",":8080","Network Address")
	flag.Parse()
	mux := http.NewServeMux()
	logError := log.New(os.Stderr, "ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)
	logInfo := log.New(os.Stdout, "INFO\t",log.Ldate|log.Ltime)
	app := &application{
		logError: logError,
		logInfo: logInfo,
	}
	mux.HandleFunc("/home",app.home)
	mux.HandleFunc("/about_us",app.aboutUs)
	mux.HandleFunc("/view",app.view)

	srv := http.Server{
		Addr: *addr,
		Handler: mux,
	}
	logInfo.Println("starting server at",*addr)
	err := srv.ListenAndServe()

	logError.Panic(err)
}
