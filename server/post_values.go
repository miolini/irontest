package server

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func (app *App) postValues(ctx *echo.Context) (err error) {
	key := ctx.Param("key")
	lockID := ctx.Param("lock_id")

	release, err := strconv.ParseBool(ctx.Query("release"))
	if err != nil {
		return err
	}

	value, err := readBody(ctx.Request())
	if err != nil {
		return err
	}

	app.Lock()
	entry, ok := app.store[key]
	app.Unlock()

	if !ok {
		ctx.NoContent(http.StatusNotFound)
		return
	}

	entry.Lock()
	defer entry.Unlock()

	if entry.lockID != lockID {
		ctx.NoContent(http.StatusUnauthorized)
		return
	}

	entry.value = value

	if release {
		entry.lockID = ""
	}

	ctx.NoContent(http.StatusNoContent)

	return
}
