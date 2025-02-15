// Handlers
package handlers

import "net/http"

// HTTP request handler
func Webhook(w http.ResponseWriter, r *http.Request) {
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
