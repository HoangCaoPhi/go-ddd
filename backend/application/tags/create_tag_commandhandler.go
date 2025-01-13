package tags

import (
	"phihc116/go-task/backend/application/auth"
	"phihc116/go-task/backend/domain/tags"
	"phihc116/go-task/backend/shared"

	"github.com/google/uuid"
)

type CreateTagCommandHandler struct {
	TagRepository   tags.TagRepo
	IdentityService auth.IdentityService
}

func NewCreateTagCommandHandler(
	tagRepository tags.TagRepo,
	identityService auth.IdentityService) *CreateTagCommandHandler {
	return &CreateTagCommandHandler{tagRepository, identityService}
}

func (h *CreateTagCommandHandler) Handler(command CreateTagCommand) shared.Result[uuid.UUID] {
	tag := tags.CreateTag(command.Name, h.IdentityService.GetUserId())
	result := h.TagRepository.CreateTag(tag)
	return result
}
