package tags

import (
	"phihc116/go-task/backend/shared"

	"github.com/google/uuid"
)

type TagRepo interface {
	CreateTag(tag *Tag) shared.Result[uuid.UUID]
}
