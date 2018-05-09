package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func getArticleInfo(c *gin.Context) {
	c.String(http.StatusOK, "Hello World!")
}

func RegisterArticleHandlers(r *gin.Engine) {
	ar := r.Group("/articles")
	ar.GET("/:articleID", getArticleInfo)
}
