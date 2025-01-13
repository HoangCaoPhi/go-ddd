package presentation

import (
	"phihc116/go-task/backend/presentation/tags"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/tags", tags.NewCreateTag)

	return router
}
