// Yandex Alice Skill
package main

import (
	"github.com/srg-bnd/alice-skill/internal/logger"
	"github.com/srg-bnd/alice-skill/internal/server"
)

type App struct {
	server *server.Server
}

func NewApp() *App {
	return &App{
		server: server.NewServer(),
	}
}

func main() {
	parseFlags()

	if err := logger.Initialize(flagLogLevel); err != nil {
		panic(err)
	}

	if err := NewApp().server.Run(flagRunAddr); err != nil {
		panic(err)
	}
}
