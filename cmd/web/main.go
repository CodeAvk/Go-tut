package main

import (
	"fmt"
	"net/http"

	"example.com/hello_world/pkg/handlers"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s\n", portNumber)

	err := http.ListenAndServe(portNumber, nil)
	if err != nil {
		fmt.Printf("Error starting the server: %s\n", err)
	}
}
