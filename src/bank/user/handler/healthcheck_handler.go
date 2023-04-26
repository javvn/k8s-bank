package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type HealthCheckHandler struct{}

func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

func (handler *HealthCheckHandler) HealthCheck() echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, nil)
	}
}
