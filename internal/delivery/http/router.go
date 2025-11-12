package http

import ( 
	"github.com/gin-gonic/gin"
)

func NewRouter(handler *UrlHandler) * gin.Engine {
	router := gin.Default()

	api := router.Group("/api")
	{
		api.POST("/urls", handler.CreateUrl)
	}

	return router
}