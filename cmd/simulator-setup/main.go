package main

import (
	"deeplink/internal/api"
	"fmt"
	"net/http"
)

func main() {
	


	fmt.Println("Starting web server...")

	// Create a dedicated ServeMux instead of using DefaultServeMux:
	mux := http.NewServeMux()


	// API Endpoints:
	mux.HandleFunc("/api/status", api.StatusHandler)
	mux.HandleFunc("/api/setup", api.SetupHandler)
	mux.HandleFunc("/api/teardown", api.TeardownHandler)


	// Serve static assets from /static/..., mapped to web/static/:
	staticFS := http.FileServer(http.Dir("../../web/static"))
	mux.Handle("/static/", http.StripPrefix("/static/", staticFS))


	// Serve the site root:
	mux.Handle("/", http.FileServer(http.Dir("../../web")))


	// Start the HTTP server:
	port := ":8080"
	fmt.Printf("Control Panel available at http://localhost%v\n", port)

	// http.ListenAndServe bloacks the main thread and start listening for requests.
	if err := http.ListenAndServe(port, mux); err != nil {
		fmt.Println("Server failed to start: ", err)
	}
}
