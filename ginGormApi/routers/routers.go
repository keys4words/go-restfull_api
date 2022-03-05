package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/keys4words/go-restfull_api/ginGormApi/handlers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("articles", handlers.GetAllArticles)
		apiGroup.POST("article", handlers.PostNewArticle)
		apiGroup.GET("article/:id", handlers.GetArticleById)
		apiGroup.PUT("article/:id", handlers.UpdateArticleById)
		apiGroup.DELETE("article/:id", handlers.DeleteArticleById)
	}
	// apiUserGroup := router.Group("/user")
	return router
}
