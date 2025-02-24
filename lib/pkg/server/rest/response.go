package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewPayloadErrorResponse(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusBadRequest, map[string]string{
		"error": fmt.Sprintf("invalid payload: %s", msg),
	})
}

func NewCreatedResponse(ctx echo.Context, entity, resultantID string) error {
	return ctx.JSON(http.StatusCreated, map[string]any{
		"message": fmt.Sprintf("%s created successfully", entity),
		"data":    map[string]string{"id": resultantID},
	})
}

func NewConflictErrorResponse(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusConflict, map[string]any{
		"message": msg,
	})
}

func NewInternalServerErrorReponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]any{
		"message": "internal server error",
	})
}
