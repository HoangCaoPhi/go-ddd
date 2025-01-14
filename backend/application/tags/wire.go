package tags

import (
	"phihc116/go-task/backend/infrastructure/auth"
	"phihc116/go-task/backend/infrastructure/persistence/repositories"

	"github.com/google/wire"
)

var CreateTagSets = wire.NewSet(
	NewCreateTagCommandHandler,
	repositories.NewTagRepoImpl,
	auth.NewIdentity,
)
