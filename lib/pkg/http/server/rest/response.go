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

func NewUnauthorizedErrorResponse(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusUnauthorized, map[string]string{
		"error": msg,
	})
}

func NewCreatedResponse(ctx echo.Context, entity, resultantID string) error {
	return ctx.JSON(http.StatusCreated, map[string]any{
		"message": fmt.Sprintf("%s created successfully", entity),
		"data":    map[string]string{"id": resultantID},
	})
}

func NewOkResponse(ctx echo.Context, data any) error {
	return ctx.JSON(http.StatusOK, data)
}

func NewConflictErrorResponse(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusConflict, map[string]any{
		"message": msg,
	})
}

func NewBadRequestResponse(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusBadRequest, map[string]any{
		"message": msg,
	})
}

func NewInternalServerErrorReponse(ctx echo.Context) error {
	return ctx.JSON(http.StatusInternalServerError, map[string]any{
		"message": "internal server error",
	})
}

func NewResourceNotFoundErrResponse(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusNotFound, msg)
}

func NewUnprocessableEntity(ctx echo.Context, msg string) error {
	return ctx.JSON(http.StatusUnprocessableEntity, map[string]any{
		"message": msg,
	})
}
