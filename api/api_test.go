package api

import (
	"encoding/json"
	"github.com/labstack/echo"
	gen "github.com/pintobikez/price-service/api/structures"
	mock "github.com/pintobikez/price-service/mocks"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

/* Test for ValidateProduct method */
func TestValidateProduct(t *testing.T) {

	m := make(map[string]int64)
	m["TESTE"] = 1

	p := new(gen.Product)
	a := new(API)
	v := a.validateProduct(p, m)
	assert.Equal(t, "is empty", v["id"])
	assert.Equal(t, "is empty", v["prices"])

	pp := new(gen.ProductPrices)
	pp.Price = 0
	p.Prices = append(p.Prices, pp)
	v = a.validateProduct(p, m)
	assert.Equal(t, "invalid value", v["price"])
	assert.Equal(t, "is not a valid value", v["channel"])

	p.Prices[0].SpecialPrice.Float64 = 10
	p.Prices[0].SpecialPrice.Valid = true
	p.Prices[0].SpecialFrom.Valid = false
	v = a.validateProduct(p, m)
	assert.Equal(t, "special date from must be supplied", v["specialFrom"])

	p.Prices[0].SpecialPrice.Float64 = 0
	p.Prices[0].SpecialPrice.Valid = true
	p.Prices[0].SpecialFrom.Valid = true
	v = a.validateProduct(p, m)
	assert.Equal(t, "special price must be supplied", v["specialPrice"])

	p.Prices[0].SpecialPrice.Float64 = 10
	p.Prices[0].SpecialPrice.Valid = true
	p.Prices[0].SpecialFrom.Valid = true
	p.Prices[0].SpecialTo.Valid = true
	p.Prices[0].SpecialFrom.Time = time.Date(2010, 1, 1, 12, 15, 30, 918273645, time.UTC)
	p.Prices[0].SpecialTo.Time = time.Date(2009, 1, 1, 12, 15, 30, 918273645, time.UTC)
	v = a.validateProduct(p, m)
	assert.Equal(t, "special date to must be after special date from", v["specialTo"])
}

/*
Test for PutProduct method
*/
type putProductProviderApi struct {
	value  string
	json   string
	result int
	code   int
	erro   string
}

var testPutProductProviderApi = []putProductProviderApi{
	{"/price/", "", http.StatusBadRequest, ErrorCodeWrongJsonFormat, ""},                                                                                                                                                                        // Incorrect url no sku
	{"/price/", `{"id":"ABCDEFGH","prices":[{"price":100.00,"specialPrice":90.00,"specialFrom":"2017-11-02T15:04:05Z","specialTo":"2017-11-22T15:04:05Z","channel":"loja1"}]}`, http.StatusInternalServerError, ErrorCodeNoChannels, "channel"}, // error retrieving channels
	{"/price/", `{"id":"ABCDEFGH","prices":[{"price":100.00,"specialPrice":90.00,"specialTo":"2017-11-22T15:04:05Z","channel":"loja1"}]}`, http.StatusBadRequest, ErrorCodeInvalidContent, ""},                                                  // error validation
	{"/price/", `{"id":"ABCDEFGH","prices":[{"price":100.00,"specialPrice":90.00,"specialFrom":"2017-11-02T15:04:05Z","specialTo":"2017-11-22T15:04:05Z","channel":"loja1"}]}`, http.StatusInternalServerError, ErrorCodeStoringContent, "put"}, // error PutProduct
	{"/price/", `{"id":"ABCDEFGH","prices":[{"price":100.00,"specialPrice":90.00,"specialFrom":"2017-11-02T15:04:05Z","specialTo":"2017-11-22T15:04:05Z","channel":"loja1"}]}`, http.StatusOK, 0, ""},                                           // Ok no Publish
	{"/price/", `{"id":"SCA","prices":[{"price":100.00,"specialPrice":90.00,"specialFrom":"2017-11-02T15:04:05Z","specialTo":"2017-11-22T15:04:05Z","channel":"loja1"}]}`, http.StatusNotFound, ErrorCodeProductNotFound, ""},                   // Product not found
	{"/price/", `{"id":"SCD","prices":[{"price":100.00,"specialPrice":90.00,"specialFrom":"2017-11-02T15:04:05Z","specialTo":"2017-11-22T15:04:05Z","channel":"loja1"}]}`, http.StatusInternalServerError, ErrorCodePublishingMessage, ""},      // Error publishing
}

func TestPutProduct(t *testing.T) {
	for _, pair := range testPutProductProviderApi {
		p := new(mock.PublisherMock)
		r := new(mock.RepositoryMock)
		a := New(r, p)

		if pair.erro == "channel" {
			r.IsErrorCh = true
		}
		if pair.erro == "put" {
			r.IsError = true
		}

		// Setup
		e := echo.New()
		e.PUT("/price/", a.PutProduct())

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("PUT", pair.value, strings.NewReader(pair.json))
		req.Header.Set("Content-Type", "application/json")
		e.ServeHTTP(rec, req)

		assert.Equal(t, rec.Code, pair.result, "Http Code doesn't match")

		if pair.result != http.StatusOK {
			erm := new(ErrResponse)
			_ = json.Unmarshal([]byte(rec.Body.String()), erm)
			assert.Equal(t, pair.code, erm.Error.Code, "ErrorCode doesn't match")
		}
	}
}

/*
Tests for GetProduct method
*/
type getProductProviderApi struct {
	value  string
	result int
}

var testGetProductProviderApi = []getProductProviderApi{
	{"/price/", http.StatusNotFound},    // url not found
	{"/price/SCA", http.StatusNotFound}, // product not found
	{"/price/SC", http.StatusOK},        // product found
}

func TestGetProduct(t *testing.T) {
	for _, pair := range testGetProductProviderApi {
		p := new(mock.PublisherMock)
		r := new(mock.RepositoryMock)
		a := New(r, p)

		// Setup
		e := echo.New()
		e.GET("/price/:id", a.GetProduct())

		rec := httptest.NewRecorder()
		req := httptest.NewRequest("GET", pair.value, strings.NewReader(""))
		req.Header.Set("Content-Type", "application/json")
		e.ServeHTTP(rec, req)
		assert.Equal(t, rec.Code, pair.result, "Http Code doesn't match")
	}
}

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
			r.IsError = true
			break
		case "pub":
			p.IsError = true
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
