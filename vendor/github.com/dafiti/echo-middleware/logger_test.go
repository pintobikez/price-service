package middleware

import (
	"bytes"
	"fmt"
	"github.com/labstack/echo"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/something", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	b := new(bytes.Buffer)

	h := LoggerWithOutput(b)(func(c echo.Context) error {
		return fmt.Errorf("Something wrong")
	})

	h(c)

	if !strings.Contains(b.String(), "500") {
		t.Errorf("Wrong error message: %s", b.String())
	}
}
