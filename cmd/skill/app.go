// app logic
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/srg-bnd/alice-skill/internal/logger"
	"github.com/srg-bnd/alice-skill/internal/models"
	"github.com/srg-bnd/alice-skill/internal/store"
	"go.uber.org/zap"
)

// `app` encapsulates all the dependencies and logic of the application
type app struct {
	store store.Store
}

// `newApp` accepts external application dependencies as input and returns a new app object
func newApp(s store.Store) *app {
	return &app{store: s}
}

func (a *app) webhook(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	if r.Method != http.MethodPost {
		logger.Log.Debug("got request with bad method", zap.String("method", r.Method))
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// deserializing the query into the model structure
	logger.Log.Debug("decoding request")
	var req models.Request
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// check that you have received a clear type of request
	if req.Request.Type != models.TypeSimpleUtterance {
		logger.Log.Debug("unsupported request type", zap.String("type", req.Request.Type))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// getting a list of messages for the current user
	messages, err := a.store.ListMessages(ctx, req.Session.User.UserID)
	if err != nil {
		logger.Log.Debug("cannot load messages for user", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// forming a text with the number of messages
	text := "There are no new messages for you."
	if len(messages) > 0 {
		text = fmt.Sprintf("For you %d new messages.", len(messages))
	}

	// first request for a new session
	if req.Session.New {
		// processing the Timezone field of the request
		tz, err := time.LoadLocation(req.Timezone)
		if err != nil {
			logger.Log.Debug("cannot parse timezone")
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// get the current time in the user's time zone
		now := time.Now().In(tz)
		hour, minute, _ := now.Clock()

		// creating a new greeting text
		text = fmt.Sprintf("The exact time %d hours, %d minutes. %s", hour, minute, text)
	}

	// filling in the response model
	resp := models.Response{
		Response: models.ResponsePayload{
			Text: text, // Alice will speak our new text.
		},
		Version: "1.0",
	}

	w.Header().Set("Content-Type", "application/json")

	// sterilizing the server response
	enc := json.NewEncoder(w)
	if err := enc.Encode(resp); err != nil {
		logger.Log.Debug("error encoding response", zap.Error(err))
		return
	}
	logger.Log.Debug("sending HTTP 200 response")
}
