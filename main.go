package main

import (
	"github.com/UmaruCMS/article-system/config"
	"github.com/UmaruCMS/article-system/http/router"
)

func release() {
	config.Release()
}

func main() {
	defer release()

	// article, err := model.NewArticle("测试文章", &model.Author{
	// 	UserID: 1,
	// }, "<p>Test Content</p>")
	// fmt.Println(article, err)
	// existedArticle := &model.Article{}
	// existedArticle.GetByUID(article.UID)
	// existedArticle.UpdateContent("<p>New Content</p>")

	r := router.DefaultRouter()
	router.RegisterHandlers(r)
	r.Run(":2335")
}
