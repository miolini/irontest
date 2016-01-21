package server

import (
	"github.com/labstack/echo"
)

func (app *App) putValues(ctx *echo.Context) (err error) {
	key := ctx.Param("key")
	value, err := readBody(ctx.Request())
	if err != nil {
		return err
	}

	app.Lock()
	entry, ok := app.store[key]
	if !ok {
		entry = newEntry()
		entry.Lock()
		app.store[key] = entry
	} else {
		entry.Lock()
	}
	app.Unlock()

	entry.value = value
	entry.lockID = generateID(32)
	app.cond.Broadcast()
	entry.Unlock()

	response := map[string]interface{}{}
	response["lock_id"] = entry.lockID
	ctx.JSON(200, response)
	return
}
