package server

import (
	"github.com/labstack/echo"
)

func (app *App) postReservations(ctx *echo.Context) (err error) {
	var entry *Entry

	key := ctx.Param("key")

	for entry == nil {
		app.Lock()
		entry = app.store[key]
		if entry == nil {
			app.cond.Wait()
			entry = app.store[key]
		}
		app.Unlock()
	}

	entry.Lock()
	response := map[string]interface{}{}
	response["value"] = entry.value
	response["lock_id"] = entry.lockID
	entry.Unlock()

	ctx.JSON(200, response)
	return
}
