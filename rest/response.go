package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const MessageOK = "200 OK"

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Detail  interface{} `json:"detail"`
}

func OK(c echo.Context, response interface{}) error {
	// Success
	return json(c, http.StatusOK, MessageOK, response)
}

func json(c echo.Context, code int, message string, response interface{}) error {
	success := false
	if code < http.StatusBadRequest {
		success = true
	}
	// Success
	return c.JSON(code, Response{
		Success: success,
		Message: message,
		Detail:  response,
	})
}

func raise(c echo.Context, code int, message string) error {
	// Success
	return json(c, code, message, nil)
}
