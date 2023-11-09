package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	fmt.Println("startup")
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) Run() {
	go func() {
		count := 1

		for {
			if a.ctx == nil {
				log.Printf("ctx is nil")
				time.Sleep(1 * time.Second)
				continue
			}

			log.Printf("emitting count: %v", count)
			runtime.EventsEmit(a.ctx, "count", count)
			time.Sleep(1 * time.Second)
			count += 1
		}
	}()
}
