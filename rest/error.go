package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	MessageBadRequest     = "400 Bad Request"
	MessageUnauthorized   = "401 Unauthorized"
	MessageForbidden      = "403 Forbidden"
	MessageNotFound       = "404 Not Found"
	MessageConflict       = "409 Conflict"
	MessageInternalServer = "500 Internal Server Error"
)

func RaiseNotFound(c echo.Context, err error) error {
	logger.Errorf("[REST][%d]: %v", http.StatusNotFound, err)
	// Success
	return raise(c, http.StatusNotFound, MessageNotFound)
}

func RaiseBadRequest(c echo.Context, err error) error {
	logger.Errorf("[REST][%d]: %v", http.StatusBadRequest, err)
	// Success
	return raise(c, http.StatusBadRequest, MessageBadRequest)
}

func RaiseUnauthorized(c echo.Context) error {
	logger.Errorf("[REST][%d]: %v", http.StatusUnauthorized, MessageUnauthorized)
	// Success
	return raise(c, http.StatusUnauthorized, MessageUnauthorized)
}

func RaiseForbidden(c echo.Context) error {
	logger.Errorf("[REST][%d]: %v", http.StatusForbidden, MessageForbidden)
	// Success
	return raise(c, http.StatusForbidden, MessageForbidden)
}

func RaiseConflict(c echo.Context, err error) error {
	logger.Errorf("[REST][%d]: %v", http.StatusConflict, err)
	// Success
	return raise(c, http.StatusConflict, MessageConflict)
}

func RaiseInternalSever(c echo.Context, err error) error {
	logger.Errorf("[REST][%d]: %v", http.StatusInternalServerError, err)
	// Success
	return raise(c, http.StatusInternalServerError, MessageInternalServer)
}
