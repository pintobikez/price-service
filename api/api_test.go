package api

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	gen "github.com/pintobikez/price-service/api/structures"
	mock "github.com/pintobikez/price-service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

/*
Tests for HealthStatus method
*/
type getHealthStatusApi struct {
	value string
	erro  string
}

var testGetHealthStatusApi = []getHealthStatusApi{
	{"/health/", "repo"}, // error in repo
	{"/health/", "pub"},  // error in publisher
	{"/health/", ""},     // all good
}

func TestHealthStatus(t *testing.T) {
	for _, pair := range testGetHealthStatusApi {
		p := new(mock.PublisherMock)
		r := new(mock.RepositoryMock)
		a := New(r, p)

		switch pair.erro {
		case "repo":
			r.Iserror = true
			break
		case "pub":
			p.Iserror = true
			break
		}

		// Setup
		e := echo.New()
		e.GET("/health/", a.HealthStatus())

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", pair.value, strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		val := new(gen.HealthStatus)
		_ = json.Unmarshal([]byte(rec.Body.String()), val)

		// Assertions
		switch pair.erro {
		case "repo":
			assert.Equal(t, StatusUnavailable, val.Repo.Status)
			break
		case "pub":
			assert.Equal(t, StatusUnavailable, val.Pub.Status)
			break
		}
	}
}
