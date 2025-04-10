package store

import (
	"context"
	"time"
)

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
}

// Message
type Message struct {
	ID      int64     // ID
	Sender  string    // Sender
	Time    time.Time // Time
	Payload string    // Text message
}
