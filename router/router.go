package router

import (
	"github.com/gin-gonic/gin"
	"github.com/lyeka/gwt/router/api/pexels"
)

// InitRouter 初始化路由
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	pexelsGroup := r.Group("/pexels")
	{
		pexelsGroup.GET("/search_photos", pexels.SearchPhotos)
	}


	return r
}
