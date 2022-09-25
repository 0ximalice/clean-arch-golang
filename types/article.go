package types

import (
	"time"
)

type Article struct {
	ID        int       `json:"id,omitempty"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type ArticleUsecase interface {
	Get(title string) (Article, error)
	Store(article Article) error
	Articles() ([]Article, error)
}

type ArticleRepository interface {
	Get(title string) (Article, error)
	Store(article Article) error
	Articles() ([]Article, error)
	NumberOfArticles() int
}
