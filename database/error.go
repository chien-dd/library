package database

import "errors"

var (
	NotFoundError = errors.New("not found error")
	ReflectError  = errors.New("reflect error")
)
