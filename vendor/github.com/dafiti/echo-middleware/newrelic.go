package middleware

import (
	"fmt"
	"github.com/labstack/echo"
	nr "github.com/newrelic/go-agent"
)

// NewRelic returns a middleware that collect request data for NewRelic
func NewRelic(appName string, licenseKey string) echo.MiddlewareFunc {
	config := nr.NewConfig(appName, licenseKey)
	app, err := nr.NewApplication(config)

	if err != nil {
		panic(fmt.Errorf("New relic: %s", err))
	}

	return NewRelicWithApplication(app)
}

// NewRelicWithApplication returns a NewRelic middleware with application.
// See: `NewRelic()`.
func NewRelicWithApplication(app nr.Application) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			txn := app.StartTransaction(c.Path(), c.Response().Writer, c.Request())
			defer txn.End()

			c.Set("newrelic-txn", txn)

			err := next(c)

			if err != nil {
				txn.NoticeError(err)
			}

			return err
		}
	}
}
