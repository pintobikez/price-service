package api

import (
	"fmt"
	"github.com/labstack/echo"
	strut "github.com/pintobikez/price-service/api/structures"
	pub "github.com/pintobikez/price-service/publisher"
	repo "github.com/pintobikez/price-service/repository"
	"net/http"
)

const (
	StatusAvailable   = "Available"
	StatusUnavailable = "Unavailable"
)

type API struct {
	rp repo.Repository
	pb pub.PubSub
}

func New(rpo repo.Repository, p pub.PubSub) *API {
	return &API{rp: rpo, pb: p}
}

//HealthStatus Handler for Health Status
func (a *API) HealthStatus() echo.HandlerFunc {
	return func(c echo.Context) error {

		resp := &strut.HealthStatus{
			Pub:  &strut.HealthStatusDetail{Status: StatusAvailable, Detail: ""},
			Repo: &strut.HealthStatusDetail{Status: StatusAvailable, Detail: ""},
		}

		if err := a.pb.Health(); err != nil {
			resp.Pub.Status = StatusUnavailable
			resp.Pub.Detail = err.Error()
		}
		if err := a.rp.Health(); err != nil {
			resp.Repo.Status = StatusUnavailable
			resp.Repo.Detail = err.Error()
		}

		return c.JSON(http.StatusOK, resp)
	}
}

//GetStock Handler to GET Stock request
func (a *API) GetProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		ProductResponse, err := a.rp.FindProduct(c.Param("id"))

		if err != nil {
			return c.JSON(http.StatusNotFound, &ErrResponse{ErrContent{ErrorCodeProductNotFound, err.Error()}})
		}

		return c.JSON(http.StatusOK, ProductResponse)
	}
}

//PutStock Handler to PUT Stock request
func (a *API) PutProduct() echo.HandlerFunc {
	return func(c echo.Context) error {

		var s *strut.Product

		if err := c.Bind(&s); err != nil {
			return c.JSON(http.StatusBadRequest, &ErrResponse{ErrContent{ErrorCodeWrongJsonFormat, err.Error()}})
		}

		//load all channels from db
		chs, err := a.rp.GetChannels()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &ErrResponse{ErrContent{ErrorCodeNoChannels, err.Error()}})
		}

		// check if the json is valid
		if err := a.validateProduct(s, chs); len(err) > 0 {
			return c.JSON(http.StatusBadRequest, buildErrorResponse(err))
		}

		af, err := a.rp.PutProduct(s)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, &ErrResponse{ErrContent{ErrorCodeStoringContent, err.Error()}})
		}

		if af > 0 { //publish message
			ProductResponse, err := a.rp.FindProduct(s.ID)
			if err != nil {
				return c.JSON(http.StatusNotFound, &ErrResponse{ErrContent{ErrorCodeProductNotFound, fmt.Sprintf(ProductNotFound, s.ID)}})
			}

			if err := a.pb.Publish(ProductResponse); err != nil {
				return c.JSON(http.StatusInternalServerError, &ErrResponse{ErrContent{ErrorCodePublishingMessage, err.Error()}})
			}
		}

		return c.NoContent(http.StatusOK)
	}
}

//validateProduct Validates the consistency of the Product update/insert rquest
func (a *API) validateProduct(s *strut.Product, ch map[string]int64) map[string]string {

	ret := make(map[string]string)

	if s.ID == "" {
		ret["id"] = "is empty"
	}

	for _, el := range s.Prices {
		if el.Price <= 0 {
			ret["price"] = "invalid value"
		}
		if !el.SpecialFrom.Valid && el.SpecialPrice.Valid && el.SpecialPrice.Float64 > 0 {
			ret["specialFrom"] = "special date from must be supplied"
		}
		if el.SpecialFrom.Valid && el.SpecialPrice.Valid && el.SpecialPrice.Float64 <= 0 {
			ret["specialPrice"] = "special price must be supplied"
		}
		if el.SpecialFrom.Valid && el.SpecialPrice.Float64 > 0 && el.SpecialTo.Valid && el.SpecialFrom.Time.After(el.SpecialTo.Time) {
			ret["specialDateTo"] = "special date to must be after special date from"
		}
		cid, ok := ch[el.Channel]
		if !ok {
			ret["channel"] = "is not a valid value"
		} else {
			el.ChannelID = cid
		}
	}

	return ret
}

//buildErrorResponse builds a validation error reponse strut
func buildErrorResponse(err map[string]string) *ErrResponseValidation {

	ret := &ErrResponseValidation{Type: "validation", Errors: make([]*ErrValidation, 0, len(err))}
	i := 0

	for k, v := range err {
		ret.Errors[i] = &ErrValidation{Field: k, Message: v}
		i++
	}

	return ret
}
