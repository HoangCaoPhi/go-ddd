package repositories

import (
	"phihc116/go-task/backend/domain/tags"
	"phihc116/go-task/backend/shared"

	"github.com/google/uuid"
)

type GormTagRepo struct {
}

func NewTagRepoImpl() tags.TagRepo {
	return &GormTagRepo{}
}

func (g *GormTagRepo) CreateTag(tag *tags.Tag) shared.Result[uuid.UUID] {
	result := shared.DbContext.Create(tag)

	if result.Error != nil {
		return shared.Result[uuid.UUID]{Value: uuid.Nil, Error: result.Error}
	}

	return shared.Result[uuid.UUID]{Value: tag.TagID, Error: result.Error}
}
