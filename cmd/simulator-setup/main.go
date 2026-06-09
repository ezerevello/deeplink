package main

import (
	"fmt"
	"net/http"
	"deeplink/internal/api"
)

func main() {
	fmt.Println("Starting web server...")

	// 1) Configure FileServer for the local web.
	fs := http.FileServer(http.Dir("../../web"))

	// 2) Bing routes.
	// Route "/" will serve our static files (frontend UI):
	http.Handle("/", fs)

	// Route "api/status" will trigger the statusHandler func:
	http.HandleFunc("/api/status", api.StatusHandler)

	// 3) Start the HTTP server:
	port := ":8080"
	fmt.Printf("Control Panel available at http://localhost%v\n", port)

	// http.ListenAndServe bloacks the main thread and start listening for requests.
	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}