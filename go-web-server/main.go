package main

import (
	"fmt"
	"net/http"
)

func main() {
		//  example slice
		var greetings  []string
		greetings = append(greetings, "Hello")
		greetings = append(greetings, "world")
		greetings = append(greetings, "!")
	
	// Define a route and handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my Go web server!")
		fmt.Fprintln(w, greetings)
		fmt.Fprintln(w,len(greetings))
		fmt.Fprintln(w,cap(greetings))

	})

	var strings [3]string

	strings[0] = "First"
	strings[1] = "Second"
	strings[2] = "Third"



	// Define another route
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello from Go!")
		fmt.Fprintln(w, "Array contents:")
		fmt.Fprintln(w, strings)
	})

	// Start the server on port 8080
	fmt.Println("Server running on http://localhost:8080")
	fmt.Println(strings)
	http.ListenAndServe(":8080", nil)
}
