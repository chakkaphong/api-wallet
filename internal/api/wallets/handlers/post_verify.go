package handlers

import (
	"api-wallet/internal/api/wallets/models"
	"api-wallet/internal/entities"
	"api-wallet/internal/global/errors"
	"api-wallet/internal/global/responses"
	"context"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (h *WalletHandler) PostVerify(c echo.Context) error {
	var req models.VerifyRequest

	ctx, cancel := context.WithTimeout(c.Request().Context(), 1*time.Minute)
	defer cancel()

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequest(err.Error()))
	}

	if err := c.Validate(req); err != nil {
		return c.JSON(http.StatusBadRequest, responses.BadRequest(err.Error()))
	}

	resp, err := h.service.VerifyTransaction(ctx, req)
	if err != nil {
		appErr, ok := err.(errors.AppError)
		if ok {
			return c.JSON(appErr.Code, responses.ResponseFailed(err))
		}
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := h.buildVerifyData(resp, req.UserId)

	return c.JSON(http.StatusCreated, responses.Success(data))
}

func (h *WalletHandler) buildVerifyData(transaction *entities.Transaction, userId int) models.VerifyResponse {
	return models.VerifyResponse{
		TransactionId: transaction.TransactionID,
		UserId:        userId,
		Amount:        transaction.Amount.InexactFloat64(),
		PaymentMethod: transaction.PaymentMethod,
		Status:        transaction.Status,
		ExpiresAt:     transaction.ExpiresAt,
	}
}
