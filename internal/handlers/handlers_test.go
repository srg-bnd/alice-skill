package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWebhook(t *testing.T) {
	successBody := `{
			"response": {
					"text": "Sorry, I can't do anything yet."
			},
			"version": "1.0"
	}`

	testCases := []struct {
		method       string
		expectedCode int
		expectedBody string
	}{
		{method: http.MethodGet, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPut, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodDelete, expectedCode: http.StatusMethodNotAllowed, expectedBody: ""},
		{method: http.MethodPost, expectedCode: http.StatusOK, expectedBody: successBody},
	}

	for _, tc := range testCases {
		t.Run(tc.method, func(t *testing.T) {
			r := httptest.NewRequest(tc.method, "/", nil)
			w := httptest.NewRecorder()

			Webhook(w, r)

			assert.Equal(t, tc.expectedCode, w.Code, "The response code does not match the expected one")
			if tc.expectedBody != "" {
				assert.JSONEq(t, tc.expectedBody, w.Body.String(), "The response body does not match the expected one")
			}
		})
	}
}
