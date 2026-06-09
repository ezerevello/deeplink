package api

import (
	"fmt"
	"net/http"
)

func StatusHandler(w http.ResponseWriter, r *http.Request) {
	// Check por GET method:
	if r.Method == http.MethodGet {
		// This is what the frontend will receive
		fmt.Fprintf(w, "Deep Space Network status: Currently Offline.")
	}
}