package api

import "github.com/gin-gonic/gin"

func InitArticlesRouter(router *gin.Engine) {
	v1 := router.Group("v1/articles")
	v1.GET("/:article_id", getArticle)
	v1.GET("/", getArticles)
	v1.POST("", createArticle)
	v1.PUT("/:article_id", updateArticle)
	v1.DELETE("/:article_id", deleteArticle)
}

// 获取文章列表
func getArticles(ctx *gin.Context) {
	ctx.Param("article_id")
	//根据db查询article 然后返回
}

// 获取根据文章id获取文章
func getArticle(ctx *gin.Context) {
	ctx.Param("article_id")
	//根据db查询article 然后返回
}

// 创建文章
func createArticle(ctx *gin.Context) {
	//根据db查询article 然后返回
}

// 文章更新
func updateArticle(ctx *gin.Context) {
	//根据db查询article 然后返回
}

// 文章删除
func deleteArticle(ctx *gin.Context) {
	//根据db查询article 然后返回
}
