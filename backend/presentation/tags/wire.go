//go:build wireinject
// +build wireinject

package tags

import (
	"phihc116/go-task/backend/application/tags"

	"github.com/google/wire"
)

func InitializeCreateTagHandler() *tags.CreateTagCommandHandler {
	wire.Build(tags.CreateTagSets)
	return nil
}
