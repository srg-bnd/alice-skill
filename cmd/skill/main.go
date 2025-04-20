// Yandex Alice Skill
package main

import (
	"database/sql"
	"net/http"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/srg-bnd/alice-skill/internal/logger"
	"github.com/srg-bnd/alice-skill/internal/store/pg"
	"go.uber.org/zap"
)

// gzip-middleware for http-server
func gzipMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ow := w

		acceptEncoding := r.Header.Get("Accept-Encoding")
		supportsGzip := strings.Contains(acceptEncoding, "gzip")
		if supportsGzip {
			cw := NewCompressWriter(w)
			ow = cw
			defer cw.Close()
		}

		contentEncoding := r.Header.Get("Content-Encoding")
		sendsGzip := strings.Contains(contentEncoding, "gzip")
		if sendsGzip {
			cr, err := NewCompressReader(r.Body)
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

// logic for running the app
func run() error {
	if err := logger.Initialize(flagLogLevel); err != nil {
		return err
	}

	conn, err := sql.Open("pgx", flagDatabaseURI)
	if err != nil {
		return err
	}

	// creating an instance of the app, so far without the external dependency of the message store
	appInstance := newApp(pg.NewStore(conn))

	logger.Log.Info("Running server", zap.String("address", flagRunAddr))
	// wrap the webhook handler in middleware with logging and gzip support
	return http.ListenAndServe(flagRunAddr, logger.RequestLogger(gzipMiddleware(appInstance.webhook)))
}

func main() {
	parseFlags()

	if err := run(); err != nil {
		panic(err)
	}
}
