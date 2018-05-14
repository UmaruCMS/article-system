package api

import (
	"context"
	"time"

	"github.com/UmaruCMS/article-system/config"
	"github.com/UmaruCMS/article-system/rpc/client/auth"
	"github.com/gin-gonic/gin"
)

func GetAuthorizationHeader(c *gin.Context) string {
	rawString := c.GetHeader("Authorization")
	if len(rawString) < 8 {
		return ""
	}
	return rawString[7:]
}

func Auth(c *gin.Context) (*auth.Token, error) {
	rpc := config.RPC
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	token := &auth.Token{
		Raw: GetAuthorizationHeader(c),
	}
	return rpc.AuthClient.VerifyToken(ctx, token)
}

func DefaultHandlerFactory(defaultResp interface{}) func(*gin.Context) {
	if defaultResp == nil {
		defaultResp = gin.H{
			"ping": "pong",
		}
	}
	return func(c *gin.Context) {
		c.JSON(200, defaultResp)
	}
}
