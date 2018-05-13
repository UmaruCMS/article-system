package model

import (
	"fmt"
	"io"
	"os"

	"github.com/UmaruCMS/article-system/config"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	UID        uint64 `gorm:"unique_key"`
	Title      string
	AuthorID   uint
}

func NewArticle(title string, author *Author, content string) (*Article, error) {
	uid, err := config.UIDGenerator.NextID()
	filePath := fmt.Sprintf("%s/%d.html", config.RootPath, uid)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	article := &Article{
		UID:        uid,
		Title:      title,
		AuthorID:   author.UserID,
	}
	_, err = io.WriteString(file, content)
	if err != nil {
		article.removeContent()
		return nil, err
	}
	db := config.Database
	err = db.Create(article).Error
	if err != nil {
		article.removeContent()
		return nil, err
	}
	return article, nil
}

func (article *Article) GetContentPath() string {
	return fmt.Sprintf("%s/%d.html", config.RootPath, article.UID)
}

func (article *Article) removeContent() error {
	filePath := article.GetContentPath()
	return os.Remove(filePath)
}

func (article *Article) Content() (string, error) {
	filePath := article.GetContentPath()
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return "", err
	}
	buf := make([]byte, 1024)
	chunks := make([]byte, 1024, 1024)

	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return "", err
		}
		if n == 0 {
			break
		}
		chunks = append(chunks, buf[:n]...)
	}
	return string(chunks), nil
}

func (article *Article) GetByUID(uid uint64) (*Article, error) {
	db := config.Database
	if err := db.Where("uid = ?", uid).Find(article).Error; err != nil {
		return nil, err
	}
	return article, nil
}

func (article *Article) Delete(permanently bool) error {
	db := config.Database
	if permanently {
		return db.Delete(article).Error
	}
	return db.Unscoped().Delete(article).Error
}

func (article *Article) 