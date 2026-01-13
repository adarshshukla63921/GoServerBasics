package main

import (
	"flag"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte(`<!DOCTYPE html>
	<html lang="en">
	<head>
	  <meta charset="UTF-8" />
	  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
	  <title>My Website</title>
	</head>
	<body>
	
	  <!-- Header -->
	  <header>
		<h1>My Website</h1>
		<nav>
		  <a href="#">Home</a> |
		  <a href="#">About</a> |
		  <a href="#">Services</a> |
		  <a href="#">Contact</a>
		</nav>
	  </header>
	
	  <hr />
	
	  <!-- Hero Section -->
	  <section>
		<h2>Welcome to My Website</h2>
		<p>
		  Here we learn to use command line flags in Go, though dont use the html as i have.
		</p>
		<button>Get Started</button>
	  </section>
	
	  <hr />
	
	  <!-- Features Section -->
	  <section>
		<h2>Using 'flag'</h2>
	
		<article>
		  <h3>Step 1</h3>
		  <p>addr := flag.String(name, defaultValue, short string explaining user), this returns a pointer</p>
		</article>
	
		<article>
		  <h3>Step 2</h3>
		  <p>use flag.Parse(), this will get the value of the command line flag, and assign to your variable</p>
		</article>
	
		<article>
		  <h3>Step 4</h3>
		  <p>Pass it the server by deferencing the pointer, and use 'go run . -addr=":8080" or any other port number you like'</p>
		</article>
	  </section>
	
	  <hr />
	  <!-- Footer -->
	  <footer>
		<p>&copy; 2026 My Website. All rights reserved.</p>
	  </footer>
	
	</body>
	</html>
	`))
}

func main(){
	addr := flag.String("addr",":8080","Http server port")

	flag.Parse()
	mux := http.NewServeMux()

	mux.HandleFunc("/",home)

	err := http.ListenAndServe(*addr,mux)

	log.Fatal(err)
	
}