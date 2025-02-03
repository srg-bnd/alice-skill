// Skill
package main

import (
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

// Init server dependencies before startup
func run() error {
	return http.ListenAndServe(`:8080`, http.HandlerFunc(webhook))
}

// HTTP request handler
func webhook(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// only POST-requests
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// set headers
	w.Header().Set("Content-Type", "application/json")
	// TODO: temporary stub
	_, _ = w.Write([]byte(`
      {
        "response": {
          "text": "Sorry, I can't do anything yet."
        },
        "version": "1.0"
      }
    `))
}
