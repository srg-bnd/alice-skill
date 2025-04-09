// Handlers
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"

	"github.com/srg-bnd/alice-skill/internal/logger"
	"github.com/srg-bnd/alice-skill/internal/models"
)

// HTTP request handler
func Webhook(w http.ResponseWriter, r *http.Request) {
	// only POST-requests
	if r.Method != http.MethodPost {
		logger.Log.Debug("got request with bad method", zap.String("method", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	logger.Log.Debug("decoding request")
	var req models.Request
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	text := "There are no new messages for you."
	if req.Session.New {
		tz, err := time.LoadLocation(req.Timezone)
		if err != nil {
			logger.Log.Debug("cannot parse timezone")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		now := time.Now().In(tz)
		hour, minute, _ := now.Clock()
		text = fmt.Sprintf("The exact time %d hours, %d minutes. %s", hour, minute, text)
	}

	if req.Request.Type != models.TypeSimpleUtterance {
		logger.Log.Debug("unsupported request type", zap.String("type", req.Request.Type))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	resp := models.Response{
		Response: models.ResponsePayload{
			Text: text,
		},
		Version: "1.0",
	}

	w.Header().Set("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		logger.Log.Debug("error encoding response", zap.Error(err))
		return
	}

	logger.Log.Debug("sending HTTP 200 response")
}
