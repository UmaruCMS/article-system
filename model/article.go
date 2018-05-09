package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title   string
	Author  *Author
	content string
}

func NewArticle(title string, author *Author, content string) *Article {
	return nil
}

func (article *Article) Content() string {
	return article.content
}
