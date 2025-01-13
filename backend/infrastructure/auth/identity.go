package auth

import (
	"phihc116/go-task/backend/application/auth"

	"github.com/google/uuid"
)

type Identity struct{}

func NewIdentity() auth.IdentityService {
	return &Identity{}
}

func (i *Identity) GetUserId() uuid.UUID {
	return uuid.New()
}
