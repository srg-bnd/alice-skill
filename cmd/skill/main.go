// Yandex Alice Skill
package main

import "github.com/srg-bnd/alice-skill/internal/server"

type App struct {
	server *server.Server
}

func NewApp() *App {
	return &App{
		server: server.NewServer(),
	}
}

func main() {
	if err := NewApp().server.Run(); err != nil {
		panic(err)
	}
}
