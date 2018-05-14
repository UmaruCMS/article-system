package api

import (
	"net/http"
	"strconv"

	"github.com/UmaruCMS/article-system/controller/article"
	"github.com/UmaruCMS/article-system/controller/author"
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
	authorName := "-"
	author, err := author.GetAuthor(article.AuthorID)
	if err == nil {
		authorName = author.Name
	}
	c.JSON(http.StatusOK, &gin.H{
		"title":       article.Title,
		"author_name": authorName,
		"author_id":   article.AuthorID,
		"content":     content,
	})
}

type articleForm struct {
	Title   string `form:"title" binding:"required"`
	Content string `form:"content" binding:"required"`
}

func createArticle(c *gin.Context) {
	articleForm := &articleForm{}
	err := c.ShouldBind(articleForm)
	token, err := Auth(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"error": err.Error()})
		return
	}
	if !token.Valid {
		c.JSON(http.StatusForbidden, &gin.H{"error": "permission denied"})
	}
	article, err := article.CreateArticle(uint(token.UserId), articleForm.Title, articleForm.Content)
	authorName := "-"
	author, err := author.GetAuthor(article.AuthorID)
	if err == nil {
		authorName = author.Name
	}
	c.JSON(http.StatusCreated, &gin.H{
		"title":       article.Title,
		"author_name": authorName,
		"author_id":   token.UserId,
		"content":     articleForm.Content,
	})
}

func updateArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Not a valid integer ID")
		return
	}
	targetArticle, err := article.GetArticle(uint(articleID))
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"error": err.Error()})
		return
	}
	token, err := Auth(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"error": err.Error()})
		return
	}
	if !token.Valid || uint(token.UserId) != targetArticle.AuthorID {
		c.String(http.StatusForbidden, "not allowed")
		return
	}
	articleForm := &articleForm{}
	err = c.ShouldBind(articleForm)
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"error": err.Error()})
		return
	}
	finalArticle, err := article.UpdateArticle(targetArticle.UID, articleForm.Title, articleForm.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &gin.H{"error": err.Error()})
		return
	}
	authorName := "-"
	author, err := author.GetAuthor(finalArticle.AuthorID)
	if err == nil {
		authorName = author.Name
	}
	c.JSON(http.StatusOK, &gin.H{
		"title":       articleForm.Title,
		"author_name": authorName,
		"author_id":   finalArticle.AuthorID,
		"content":     articleForm.Content,
	})
}

func deleteArticle(c *gin.Context) {
	articleID, err := strconv.Atoi(c.Param("articleID"))
	if err != nil {
		c.String(http.StatusBadRequest, "Not a valid integer ID")
		return
	}
	targetArticle, err := article.GetArticle(uint(articleID))
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"error": err.Error()})
		return
	}
	token, err := Auth(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"error": err.Error()})
		return
	}
	if !token.Valid || uint(token.UserId) != targetArticle.AuthorID {
		c.String(http.StatusForbidden, "not allowed")
		return
	}
	err = article.DeleteArticle(uint64(articleID))
	if err != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, &gin.H{"status": "ok"})
}

func RegisterArticleHandlers(r *gin.Engine) {
	ar := r.Group("/articles")
	ar.POST("/create", createArticle)
	ar.GET("/:articleID", getArticleInfo)
	ar.PUT("/:articleID", updateArticle)
	ar.DELETE("/:articleID", deleteArticle)
}
