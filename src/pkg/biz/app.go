package biz

import (
	"context"
	"fmt"
	"parse/src/pkg/g01"
	"parse/src/pkg/monitor"
	"time"
)

type App struct {
	ctx context.Context

	fm *monitor.FlowManager[g01.G01Msg]

	out *FrontOutput[g01.G01Msg]
}

// NewApp creates a new App application struct
func NewApp() *App {
	ctx := context.Background()
	a := &App{ctx: ctx}
	a.out = NewFrontOutput[g01.G01Msg]()
	fo, _ := monitor.NewFileOutput[g01.G01Msg](fmt.Sprintf("D:\\zero-utils\\client\\wails\\parse\\src\\data\\%d.txt", time.Now().Unix()))
	a.fm = monitor.NewFlowManager(ctx, &g01.TcpMonitor{}, g01.G01Codec{}, []monitor.Output[g01.G01Msg]{a.out, fo}...)
	return a
}

// Startup is called at application Startup
func (a *App) Startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// DomReady is called after front-end resources have been loaded
func (a *App) DomReady(ctx context.Context) {
	// Add your action here
}

// BeforeClose beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) BeforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) Shutdown(ctx context.Context) {
	// Perform your teardown here
}
