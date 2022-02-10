package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"http/middleware/logs"
)

func ConfigMiddleware(e *echo.Echo) {
	e.Logger.SetLevel(log.INFO)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
}

func SetupZeroLog(e *echo.Echo) {
	logger := logs.NewZeroLogger("logs", "trace.grpc")
	// Middleware
	logConfig := logs.ZeroLogConfig{
		Logger: logger,
		FieldMap: map[string]string{
			"uri":    "@uri",
			"host":   "@host",
			"method": "@method",
			"status": "@status",
		},
	}
	e.Use(logs.ZeroLogWithConfig(logConfig))
}
