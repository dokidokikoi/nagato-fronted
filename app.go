package main

import (
	"context"
	"fmt"
	initTask "nagato/internal/init"
	"nagato/internal/service"
)

// App struct
type App struct {
	ctx     context.Context
	service service.Service
}

// NewApp creates a new App application struct
func NewApp() *App {
	initTask.Init("./application.yml")
	return &App{service: service.NewService()}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	_, err := a.service.Conf().GetConfig(ctx)
	if err != nil {
		panic(err)
	}
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
