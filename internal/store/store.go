package store

import (
	"context"
	"errors"
	"time"
)

// ErrConflict indicates a data conflict in the storage.
var ErrConflict = errors.New("data conflict")

// Store
type Store interface {
	// FindRecipient
	FindRecipient(ctx context.Context, username string) (userID string, err error)
	// ListMessages
	ListMessages(ctx context.Context, userID string) ([]Message, error)
	// GetMessage
	GetMessage(ctx context.Context, id int64) (*Message, error)
	// SaveMessage
	SaveMessage(ctx context.Context, userID string, msg Message) error
	// RegisterUser registers a new user
	RegisterUser(ctx context.Context, userID, username string) error
}

// Message
type Message struct {
	ID      int64     // ID
	Sender  string    // Sender
	Time    time.Time // Time
	Payload string    // Text message
}
