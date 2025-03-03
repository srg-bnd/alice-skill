// Server
package server

import (
	"net/http"

	"github.com/srg-bnd/alice-skill/internal/handlers"
	"github.com/srg-bnd/alice-skill/internal/logger"
	"go.uber.org/zap"
)

const defaultHost = `:8080`

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

// Init & run server
func (s *Server) Run(addr string) error {

	logger.Log.Info("Running server", zap.String("address", addr))
	return http.ListenAndServe(GetAddr(addr), logger.RequestLogger(handlers.Webhook))
}

func GetAddr(addr string) string {
	if len(addr) != 0 {
		return addr
	} else {
		return defaultHost
	}
}
