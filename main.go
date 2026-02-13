package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// We use an environment variable for the port to make it DevOps-friendly
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Go DevOps Project</h1><p>Status: Running Successfully</p>")
	})

	fmt.Printf("Server starting on http://localhost:%s\n", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
