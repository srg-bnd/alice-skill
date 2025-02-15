// Server
package server

import (
	"net/http"

	"github.com/srg-bnd/alice-skill/internal/handlers"
)

const defaultHost = `:8080`

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

// Init & run server
func (s *Server) Run() error {
	return http.ListenAndServe(defaultHost, http.HandlerFunc(handlers.Webhook))
}
