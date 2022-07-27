package model

import "github.com/go-programming-tour-book/blog-service/pkg/app"

// 文章
type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageUrl string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

type ArticleSwagger struct {
	List []*Article
	Page *app.Pager
}

func (a Article) TableName() string {
	return "blog_article"
}
