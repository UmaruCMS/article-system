package api

import (
	"net/http"
	"strconv"

	"github.com/UmaruCMS/article-system/controller/article"
	"github.com/gin-gonic/gin"
)

func getArticleInfo(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Not a valid integer ID")
		return
	}
	article, err := article.GetArticle(uint(articleID))
	if err != nil {
		c.String(http.StatusNotFound, "Not Found")
		return
	}
	content, err := article.Content()
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, &gin.H{
		"title":       article.Title,
		"author_name": "UNKNOWN", // TODO: call rpc to get author name
		"author_id":   article.AuthorID,
		"content":     content,
	})
}

func RegisterArticleHandlers(r *gin.Engine) {
	ar := r.Group("/articles")
	ar.GET("/:articleID", getArticleInfo)
}
