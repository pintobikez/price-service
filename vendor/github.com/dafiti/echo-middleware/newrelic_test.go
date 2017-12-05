package middleware

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRelicWhenRaiseAError(t *testing.T) {
	defer func() {
		err := recover().(error)

		if err.Error() != "New relic: license length is not 40" {
			t.Fatalf("Wrong panic error: %s", err.Error())
		}
	}()

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewRelic("ok", "ok")(func(c echo.Context) error {
		return c.String(http.StatusTemporaryRedirect, "test")
	})

	h(c)
}

func TestNewRelicWithApplication(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/something", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	app := new(Application)

	h := NewRelicWithApplication(app)(func(c echo.Context) error {
		return fmt.Errorf("Something wrong")
	})

	h(c)
}
