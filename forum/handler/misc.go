package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog/log"
)

func UserIDFromToken(c echo.Context) uint {
	id, ok := c.Get("user").(uint)
	if !ok {
		return 0
	}
	return id
}
