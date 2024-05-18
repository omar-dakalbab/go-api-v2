package api

import (
	"fmt"
	"net/http"
)

func HandleUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Your GET /users logic
	fmt.Fprintln(w, "Fetching users...")
}
