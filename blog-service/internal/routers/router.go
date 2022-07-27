package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-programming-tour-book/blog-service/docs"
	"github.com/go-programming-tour-book/blog-service/internal/middleware"
	v1 "github.com/go-programming-tour-book/blog-service/internal/routers/api/v1"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// i18n
	r.Use(middleware.Translations())
	// register Swagger API
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	tag := v1.NewTag()
	article := v1.NewArticle()
	apiV1 := r.Group("/api/v1")

	{
		// 创建标签
		apiV1.POST("/tags", tag.Create)
		// 删除指定标签
		apiV1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiV1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiV1.GET("/tags", tag.List)

		// 创建文章
		apiV1.POST("/articles", article.Create)
		// 删除指定文章
		apiV1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiV1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiV1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiV1.GET("/articles", article.List)
	}

	return r
}
