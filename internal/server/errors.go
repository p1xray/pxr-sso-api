package server

import "errors"

var (
	ErrInvalidIdParam    = errors.New("invalid id param")
	ErrInvalidInputBody  = errors.New("invalid input body")
	ErrInvalidInputQuery = errors.New("invalid input query")
)
