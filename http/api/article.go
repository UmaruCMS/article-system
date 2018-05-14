package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/UmaruCMS/article-system/config"
	"github.com/UmaruCMS/article-system/controller/article"
	"github.com/UmaruCMS/article-system/rpc/client/user"
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
	author := &user.UserInfo{
		Name: "-",
		Id:   uint32(article.AuthorID),
	}
	rpc := config.RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	fetchedAuthor, err := rpc.UserClient.GetUserInfoByID(ctx, author)
	if err == nil {
		author = fetchedAuthor
	} else {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &gin.H{
		"title":       article.Title,
		"author_name": author.Name,
		"author_id":   article.AuthorID,
		"content":     content,
	})
}

func RegisterArticleHandlers(r *gin.Engine) {
	ar := r.Group("/articles")
	ar.GET("/:articleID", getArticleInfo)
}
