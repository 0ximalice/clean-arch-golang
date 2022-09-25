package usecases

import (
	"clean-arch-golang/types"
	"time"
)

type articleUsecase struct {
	repo types.ArticleRepository
}

func NewArticleUsecase(repo types.ArticleRepository) types.ArticleUsecase {
	return &articleUsecase{
		repo: repo,
	}
}

func (l *articleUsecase) Get(author string) (types.Article, error) {
	// Do other business logic here
	return l.repo.Get(author)
}

func (l *articleUsecase) Articles() ([]types.Article, error) {
	// Do other business logic here
	return l.repo.Articles()
}

func (l *articleUsecase) Store(article types.Article) error {
	article.ID = l.repo.NumberOfArticles() + 1
	article.CreatedAt = time.Now()
	// Do other business logic here
	return l.repo.Store(article)
}
