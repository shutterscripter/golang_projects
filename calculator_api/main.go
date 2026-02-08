package main

import (
	"fmt"
	"handlers/handlers"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// Routes
	mux.HandleFunc("/health", handleHealth)
	mux.HandleFunc("POST /add", handlers.AddHandler)
	mux.HandleFunc("POST /sub", handlers.SubHandler)
	mux.HandleFunc("POST /mul", handlers.MultiplyHandler)
	mux.HandleFunc("POST /div", handlers.DivideHandler)

	fmt.Println("Listening on port 3008")
	http.ListenAndServe(":3008", mux)

}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Health is okay!")
}
