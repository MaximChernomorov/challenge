package api

import (
	"encoding/json"
	"io"
	"net"
	"net/http"

	"github.com/friendsofgo/errors"
	"github.com/labstack/echo/v4"
)

type ipLocationResponse struct {
	ErrorResponse
	Country     string `json:"country,omitempty"`
	City        string `json:"city,omitempty"`
	Coordinates struct {
		Longitude float64 `json:"longitude,omitempty"`
		Latitude  float64 `json:"latitude,omitempty"`
	} `json:"coordinates,omitempty"`
}

type ipLocationRequest struct {
	IPAddress string `json:"ip_address"`
}

// LocateIP echo http handler
func (api *API) LocateIP(c echo.Context) error {
	response := ipLocationResponse{}
	request, err := readIPLocationRequest(c.Request().Body)
	if err != nil {
		response.Error = err.Error()
		return c.JSON(http.StatusBadRequest, response)
	}
	geoLocation, err := api.repo.LocateIP(c.Request().Context(), request.IPAddress)
	if err != nil {
		response.Error = "location not found"
		return c.JSON(http.StatusNotFound, response)
	}
	response.City = geoLocation.City.String
	response.Country = geoLocation.Country.String
	response.Coordinates.Latitude = geoLocation.Coordinates.Y
	response.Coordinates.Longitude = geoLocation.Coordinates.X

	return c.JSON(http.StatusOK, response)
}

func readIPLocationRequest(rc io.ReadCloser) (ipLocationRequest, error) {
	request := ipLocationRequest{}
	err := json.NewDecoder(rc).Decode(&request)
	if err != nil {
		return request, errors.Wrap(err, "invalid json")
	}
	if net.ParseIP(request.IPAddress) == nil {
		return request, errors.New("invalid IP address")
	}

	return request, nil
}
