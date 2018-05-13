package article

import "github.com/UmaruCMS/article-system/model"

func GetArticle(articleID uint) (*model.Article, error) {
	article := &model.Article{}
	article, err := article.GetByUID(uint64(articleID))
	if err != nil {
		return nil, err
	}
	return article, nil
}
