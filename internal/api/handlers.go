package api

import (
	"fmt"
	"net/http"
	"os/exec"
)

// StatusHandler returns the current status of deeplink (up/down)
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		// This is what the frontend will receive
		fmt.Fprintf(w, "Deep Space Network status: Currently Offline.")
	}
}

// SetupHandler initializes the namespaces
func SetupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed, use POST.", http.StatusMethodNotAllowed)
		return
	}

	// 1) Define the command
	cmd := exec.Command("/bin/bash", "../../scripts/env-up.sh")


	// 2) Execute the script and wait for its completion:
	err := cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("An error ocurred executing env-up.sh: %v", err), http.StatusInternalServerError)
		return
	}

	
	// 3) If everything goes well:
	fmt.Fprintf(w, "Namespaces created.")
}

func TeardownHandler(w http.ResponseWriter, r *http.Request) {
	// Same thing as above, but for Teardown
	if r.Method != http.MethodPost {
		http.Error(w, "Method now allowed, use POST.", http.StatusMethodNotAllowed)
		return
	}

	cmd := exec.Command("/bin/bash", "../../scripts/env-down.sh")

	err := cmd.Run()
	if err != nil {
		http.Error(w, fmt.Sprintf("An error ocurred executing env-down.sh: %v", err), http.StatusInternalServerError)
	}

	fmt.Fprintf(w, "Namespaces deleted.")
}