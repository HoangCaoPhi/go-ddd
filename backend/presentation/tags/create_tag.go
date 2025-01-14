package tags

import (
	"phihc116/go-task/backend/application/tags"
	"phihc116/go-task/backend/shared"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateTag struct {
	TagId uuid.UUID `json:"tagId"`
}

func NewCreateTag(ctx *gin.Context) {
	var createTagRequest CreateTagRequest

	if err := ctx.ShouldBindJSON(&createTagRequest); err != nil {
		shared.WriteValidateProblemDetails(ctx, err)
		return
	}

	command := tags.CreateTagCommand(createTagRequest)

	handler := InitializeCreateTagHandler()

	result := handler.Handler(command)

	result.Unwrap(ctx)
}
