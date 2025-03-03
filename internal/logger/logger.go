// Singleton logger
package logger

import (
	"net/http"

	"go.uber.org/zap"
)

// The Log will be available to the entire code as a singleton.
// No skill code other than the Initialize function should modify this variable.
// By default, the no-op logger is installed, which does not output any messages.
var Log *zap.Logger = zap.NewNop()

// Initialize initializes the logger singleton with the required logging level.
func Initialize(level string) error {
	// convert the text logging level to zap.AtomicLevel
	lvl, err := zap.ParseAtomicLevel(level)
	if err != nil {
		return err
	}
	// create a new logger configuration
	cfg := zap.NewProductionConfig()
	// устанавливаем уровень
	cfg.Level = lvl
	// create a logger based on the configuration
	zl, err := cfg.Build()
	if err != nil {
		return err
	}
	// set the singleton
	Log = zl
	return nil
}

// RequestLogger is a middleware logger for incoming HTTP requests.
func RequestLogger(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		Log.Debug("got incoming HTTP request",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
		)
		h(w, r)
	})
}
