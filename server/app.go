package server

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"log"
	"net/http"
	"sync"
)

const (
	LockIDLength = 16
)

type App struct {
	sync.Mutex
	cond   *sync.Cond
	store  map[string]*Entry
	router *echo.Echo
}

func NewApp() *App {
	app := &App{}
	app.cond = sync.NewCond(app)
	app.store = make(map[string]*Entry)
	app.router = echo.New()
	app.router.Use(mw.Logger(), mw.Recover())
	app.router.Post("/reservations/:key", app.postReservations)
	app.router.Post("/values/:key/:lock_id", app.postValues)
	app.router.Put("/values/:key", app.putValues)
	return app
}

func (app *App) Run(listenAddr string) error {
	log.Printf("http listen on %s", listenAddr)
	return http.ListenAndServe(listenAddr, app.router)
}
