// Handlers
package handlers

import (
	"net/http"

	"github.com/srg-bnd/alice-skill/internal/logger"
	"go.uber.org/zap"
)

// HTTP request handler
func Webhook(w http.ResponseWriter, r *http.Request) {
	// only POST-requests
	if r.Method != http.MethodPost {
		logger.Log.Debug("got request with bad method", zap.String("method", r.Method))
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
	logger.Log.Debug("sending HTTP 200 response")
}
