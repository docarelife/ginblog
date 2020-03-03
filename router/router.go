package router

import (
	v1 "blog/api/v1"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由引擎
func InitRouter() *gin.Engine {

	// 新建路由引擎
	r:=gin.New()

	// 全局中间件
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// 装在路由
	// 获取文章列表
	r.GET("/article",v1.GetArticleList)
	r.GET("/article/:id",v1.GetArticle)
	r.POST("/article",v1.AddArticle)
	r.PUT("/article/:id",v1.UpdateArticle)
	r.DELETE("/article/:id",v1.DeleteArticle)

	r.POST("/praise/:what/:id",v1.Praise)

	return r
}
