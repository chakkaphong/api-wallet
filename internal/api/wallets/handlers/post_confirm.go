package handlers

import (
	"api-wallet/internal/api/wallets/models"
	"api-wallet/internal/global/errors"
	"api-wallet/internal/global/responses"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *WalletHandler) PostConfirm(c echo.Context) error {
	var req models.ConfirmTransactionRequest

	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequest(err.Error()))
	}

	if err := c.Validate(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequest(err.Error()))
	}

	resp, err := h.service.ConfirmTransaction(ctx, req)
	if err != nil {
		appErr, ok := err.(errors.AppError)
		if ok {
			return c.JSON(appErr.Code, responses.ResponseFailed(err))
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, responses.Success(resp))
}
