package middleware

import (
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"io"
)

func LoggerWithOutput(w io.Writer) echo.MiddlewareFunc {
	config := mw.DefaultLoggerConfig
	config.Output = w

	return mw.LoggerWithConfig(config)
}
