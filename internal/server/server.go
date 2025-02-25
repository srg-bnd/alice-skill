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
func (s *Server) Run(addr string) error {
	return http.ListenAndServe(GetAddr(addr), http.HandlerFunc(handlers.Webhook))
}

func GetAddr(addr string) string {
	if len(addr) != 0 {
		return addr
	} else {
		return defaultHost
	}
}
