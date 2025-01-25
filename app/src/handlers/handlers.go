package handlers

import (
	"database_anonymizer/app/src/cache"
	"database_anonymizer/app/src/interfaces"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	cache cache.CacheManager
}

func NewHandler(cache cache.CacheManager) interfaces.HandlersInterface {
	return Handler{
		cache: cache,
	}
}

func (h Handler) RegisterRoutes(e *echo.Echo) {
	e.GET("/", h.Ping)

	anonymizationRequests := e.Group("anonymization-requests")
	anonymizationRequests.POST("", h.PostAnonymizationRequest)
	anonymizationRequests.GET("/:polling_key", h.GetAnonymizationStatus)
}
