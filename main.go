package main

import (
	"fmt"

	"github.com/UmaruCMS/article-system/config"
	"github.com/UmaruCMS/article-system/model"
)

func release() {
	config.Release()
}

func main() {
	defer release()
	article, err := model.NewArticle("测试文章", &model.Author{
		UserID: 1,
	}, "<p>Test Content</p>")
	fmt.Println(article, err)
}
