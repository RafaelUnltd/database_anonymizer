package handlers

import (
	"database_anonymizer/app/src/structs"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) GetAnonymizationStatus(c echo.Context) error {
	ctx := c.Request().Context()
	pollingKey := c.Param("polling_key")

	status, err := h.cache.ReadPollingStatus(ctx, pollingKey)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			structs.HttpErrorResponse{Message: err.Error(), Tag: "ERROR_READING_POLLING_STATUS"},
		)
	}

	return c.JSON(
		http.StatusOK,
		structs.HttpDataResponse{Data: status},
	)
}
