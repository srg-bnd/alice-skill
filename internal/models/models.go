// Models
package models

const (
	TypeSimpleUtterance = "SimpleUtterance"
)

// https://yandex.ru/dev/dialogs/alice/doc/request.html
type Request struct {
	Timezone string          `json:"timezone"`
	Request  SimpleUtterance `json:"request"`
	Session  Session         `json:"session"`
	Version  string          `json:"version"`
}

// Session
type Session struct {
	New  bool `json:"new"`
	User User
}

// User
type User struct {
	UserID string
}

// Command from Request
type SimpleUtterance struct {
	Type    string `json:"type"`
	Command string `json:"command"`
}

// https://yandex.ru/dev/dialogs/alice/doc/response.html
type Response struct {
	Response ResponsePayload `json:"response"`
	Version  string          `json:"version"`
}

// ResponsePayload (for the sound)
type ResponsePayload struct {
	Text string `json:"text"`
}
