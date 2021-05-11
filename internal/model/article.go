package model

import (
	"github.com/Apriil15/blog-server/pkg/app"
	"gorm.io/gorm"
)

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

// ArticleSwagger is a struct for swagger page.
type ArticleSwagger struct {
	List  []*Article
	Pager *app.Pager
}

func (a *Article) TableName() string {
	return "blog_article"
}

func (a *Article) Create(db *gorm.DB) error {

	return db.Create(&a).Error
}

// Delete an article
func (a *Article) Delete(db *gorm.DB) error {
	return db.Where("is_del = ?", 0).Delete(&a).Error
}
