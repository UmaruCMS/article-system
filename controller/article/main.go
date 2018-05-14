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

func CreateArticle(authorID uint, title string, content string) (*model.Article, error) {
	article, err := model.NewArticle(title, &model.Author{
		UserID: authorID,
	}, content)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func UpdateArticle(articleID uint64, title string, content string) (*model.Article, error) {
	article := &model.Article{}
	article, err := article.GetByUID(articleID)
	if err != nil {
		return nil, err
	}
	if err := article.UpdateTitle(title); err != nil {
		return nil, err
	}
	if err := article.UpdateContent(content); err != nil {
		return nil, err
	}
	return article, nil
}

func DeleteArticle(articleID uint64) error {
	article := &model.Article{}
	article, err := article.GetByUID(articleID)
	if err != nil {
		return err
	}
	return article.Delete(false)
}
