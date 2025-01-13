package shared

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProblemDetails struct {
	Type     string `json:"type,omitempty"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   any    `json:"detail,omitempty"`
	Instance string `json:"instance,omitempty"`
}

func WriteValidateProblemDetails(c *gin.Context, err error) {
	problem := ProblemDetails{
		Type:     "Error",
		Title:    "Server Error",
		Status:   http.StatusInternalServerError,
		Detail:   err.Error(),
		Instance: c.Request.URL.Path,
	}

	// Add Logger

	c.JSON(http.StatusInternalServerError, problem)
}

func WriteProblemDetails(c *gin.Context, status int, title string, detail any) {
	problem := ProblemDetails{
		Type:     "Error",
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: c.Request.URL.Path,
	}

	c.JSON(status, problem)
}
