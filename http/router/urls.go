package router

import (
	"github.com/UmaruCMS/article-system/http/api"
	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	api.RegisterArticleHandlers(r)
}
