package model

import "github.com/UmaruCMS/article-system/config"

func init() {
	db := config.Database
	db.AutoMigrate(&Article{})
}
