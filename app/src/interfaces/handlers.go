package interfaces

import "github.com/labstack/echo/v4"

type HandlersInterface interface {
	RegisterRoutes(e *echo.Echo)
	Ping(c echo.Context) error
	PostAnonymizationRequest(c echo.Context) error
	GetAnonymizationStatus(c echo.Context) error
}
