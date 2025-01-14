package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AppError struct {
	Code        string
	Description string
	Type        string
}

func (a AppError) Error() string {
	return a.Description
}

type Result[T any] struct {
	Value T
	Error error
}

func (r Result[T]) IsSuccess() bool {
	return r.Error == nil
}

func (r Result[T]) IsFailed() bool {
	return r.Error != nil
}

func (r Result[T]) Unwrap(c *gin.Context) {
	if r.IsFailed() {
		statusCode := http.StatusInternalServerError
		errorMessage := "One or more validation errors has occurred"

		if appErr, ok := r.Error.(AppError); ok {
			statusCode = http.StatusBadRequest
			errorMessage = appErr.Description
		}

		WriteProblemDetails(c, statusCode, r.Error.Error(), errorMessage)
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": r.Value})
}
