// app logic
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
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

	logger.Log.Debug("decoding request")
	var req models.Request
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&req); err != nil {
		logger.Log.Debug("cannot decode request JSON body", zap.Error(err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if req.Request.Type != models.TypeSimpleUtterance {
		logger.Log.Debug("unsupported request type", zap.String("type", req.Request.Type))
		w.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	// text of skill-response
	var text string

	switch true {
	// user says to send message
	case strings.HasPrefix(req.Request.Command, "Send it"):
		username, message := "", "" // parseSendCommand(req.Request.Command)

		recipientID, err := a.store.FindRecipient(ctx, username)
		if err != nil {
			logger.Log.Debug("cannot find recipient by username", zap.String("username", username), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = a.store.SaveMessage(ctx, recipientID, store.Message{
			Sender:  req.Session.User.UserID,
			Time:    time.Now(),
			Payload: message,
		})
		if err != nil {
			logger.Log.Debug("cannot save message", zap.String("recipient", recipientID), zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		text = "The message was sent successfully"

	case strings.HasPrefix(req.Request.Command, "Прочитай"):
		messageIndex := 0 // parseReadCommand(req.Request.Command)

		messages, err := a.store.ListMessages(ctx, req.Session.User.UserID)
		if err != nil {
			logger.Log.Debug("cannot load messages for user", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		text = "There are no new messages for you."
		if len(messages) < messageIndex {
			text = "There is no such message."
		} else {
			messageID := messages[messageIndex].ID
			message, err := a.store.GetMessage(ctx, messageID)
			if err != nil {
				logger.Log.Debug("cannot load message", zap.Int64("id", messageID), zap.Error(err))
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			text = fmt.Sprintf("Message from %s, sent %s: %s", message.Sender, message.Time, message.Payload)
		}

	default:
		messages, err := a.store.ListMessages(ctx, req.Session.User.UserID)
		if err != nil {
			logger.Log.Debug("cannot load messages for user", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		text = "There are no new messages for you."
		if len(messages) > 0 {
			text = fmt.Sprintf("For you %d new messages.", len(messages))
		}

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
	}

	resp := models.Response{
		Response: models.ResponsePayload{
			Text: text, // Alica says text
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
