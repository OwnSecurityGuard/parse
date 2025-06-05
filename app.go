package main

import (
	"context"
	"fmt"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

type Playback struct {
	Req  []*Tmp
	Resp []*Tmp
}
type Tmp struct {
	Timestamp int64
	Uri       string
	Msg       string
}

type Param struct {
	Index   int
	SrcPath string
	TarPath string
	TarUri  string
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}
func (a *App) A() *Playback {

	return &Playback{
		Req: []*Tmp{
			&Tmp{
				Timestamp: time.Now().Unix(),
				Uri:       "AAAA",
				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix(),
				Uri:       "AAAA2",
				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix(),
				Uri:       "AAAA3",
				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix(),
				Uri:       "AAAA4",
				Msg:       `{"uri" :"AA2AA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix(),
				Uri:       "AAA4A",
				Msg:       `{"uri" :"A4AAA", "DADA": "SDASDASDAS"}`,
			},
		},
		Resp: []*Tmp{
			&Tmp{
				Timestamp: time.Now().Unix() - 5,
				Uri:       "AAAA",
				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix() - 7,
				Uri:       "AAAA2",
				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix() + 4,
				Uri:       "AAAA3",
				Msg:       `{"uri" :"AAAA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix() + 7,
				Uri:       "AAAA4",
				Msg:       `{"uri" :"AA2AA", "DADA": "SDASDASDAS"}`,
			},
			{
				Timestamp: time.Now().Unix() + 8,
				Uri:       "AAA4A",
				Msg:       `{"uri" :"A4AAA", "DADA": "SDASDASDAS"}`,
			},
		},
	}
}
