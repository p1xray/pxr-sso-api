package request

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	ErrInvalidInputQuery = errors.New("invalid input query")
	ErrInvalidInputForm  = errors.New("invalid input form")
)

func FromQuery[T any](c *gin.Context) (T, error) {
	var inp T
	if err := c.BindQuery(&inp); err != nil {
		return inp, fmt.Errorf("%w: %w", ErrInvalidInputQuery, err)
	}

	return inp, nil
}

func FromForm[T any](c *gin.Context) (T, error) {
	var inp T
	if err := c.Bind(&inp); err != nil {
		return inp, fmt.Errorf("%w: %w", ErrInvalidInputForm, err)
	}

	return inp, nil
}
