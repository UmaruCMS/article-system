package author

import "github.com/UmaruCMS/article-system/model"

func GetAuthor(authorID uint) (*model.Author, error) {
	author := &model.Author{}
	return author.GetByUserID(authorID)
}
