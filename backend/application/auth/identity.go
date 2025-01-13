package auth

import "github.com/google/uuid"

type IdentityService interface {
	GetUserId() uuid.UUID
}
