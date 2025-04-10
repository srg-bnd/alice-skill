// Server
package server

import (
	"net/http"
	"strings"

	"github.com/srg-bnd/alice-skill/internal/gzip"
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
	return http.ListenAndServe(addr, logger.RequestLogger(gzipMiddleware(handlers.Webhook)))
}

func gzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ow := w

		acceptEncoding := r.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			cw := gzip.NewCompressWriter(w)
			ow = cw
			defer cw.Close()
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			cr, err := gzip.NewCompressReader(r.Body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			r.Body = cr
			defer cr.Close()
		}

		h.ServeHTTP(ow, r)
	}
}
