package main

import (
	"log"
	"net/http"
)

// Define a function to handle the home page
func home(w http.ResponseWriter, r *http.Request) {
	// Write the response back to the client (HTML: Hello from Snippetbox!)
	w.Write([]byte("Hello from Snippetbox!"))
}

// Define a function to display a specific snippet
func snipperView(w http.ResponseWriter, r *http.Request) {
	// Write the response back to the client (Display a speific snippet...)
	w.Write([]byte("Display a specific snippet..."))
}

// Define a function for creating a new snippet
func snippetCreate(w http.ResponseWriter, r *http.Request) {
	// Write the response back to the client (Display a form for creating a new snippet...)
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// Create an HTTP router
	mux := http.NewServeMux()

	// Define routes
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snipperView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	// Print the available routes
	log.Print("Starting server on port 4000...")

	// Start listening for incoming requests
	err := http.ListenAndServe(":4000", mux) // Listen on port 4000
	log.Fatal(err)                           // Fatal error if there's an issue setting up the server
}
