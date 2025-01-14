package webapi

import (
	"phihc116/go-task/backend/infrastructure/persistence"
	"phihc116/go-task/backend/presentation"
	"phihc116/go-task/backend/webapi/configs"
)

func Run() {
	configs.LoadConfig()

	persistence.InitDatabase()
	router := presentation.NewRouter()

	router.Run()
}
