package handlers

import (
	"context"
	"database_anonymizer/app/src/common"
	"database_anonymizer/app/src/services"
	"database_anonymizer/app/src/structs"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h Handler) PostAnonymizationRequest(c echo.Context) error {
	ctx := c.Request().Context()

	var request structs.AnonymizationRequest

	// Lê os dados do body da requisição e preenche a variável request com eles
	if err := c.Bind(&request); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			structs.HttpErrorResponse{Message: err.Error(), Tag: "ERROR_BINDING_REQUEST"},
		)
	}

	// Cria um novo polling key para armazenar o status da execução da anonimização
	pollingKey := common.NewPollingKey()
	err := h.cache.CreatePollingStatus(ctx, pollingKey)
	if err != nil {
		return c.JSON(
			http.StatusInternalServerError,
			structs.HttpErrorResponse{Message: err.Error(), Tag: "ERROR_CREATING_POLLING_STATUS"},
		)
	}

	// Inicia a execução da anonimização em uma goroutine separada
	go func() {
		services, err := services.NewService(request.InputConnectionInfo, request.OutputConnectionInfo, request.TableNames(), h.cache)
		if err != nil {
			fmt.Println(err)
		}

		// Valida a existência das tabelas e colunas passadas na requisição
		if err := services.ValidateRules(request); err != nil {
			fmt.Println(err)
		}

		err = services.Anonymize(context.Background(), request, pollingKey)
		if err != nil {
			pollingStatus, errReadChache := h.cache.ReadPollingStatus(ctx, pollingKey)
			if errReadChache != nil {
				fmt.Println(errReadChache)
			}

			pollingStatus.Status = structs.StatusError
			pollingStatus.Finished = true

			errWriteCache := h.cache.UpdatePollingStatus(ctx, pollingKey, pollingStatus)
			if errWriteCache != nil {
				fmt.Println(errWriteCache)
			}

		}
	}()

	return c.JSON(http.StatusOK, structs.HttpDataResponse{Data: pollingKey})
}
