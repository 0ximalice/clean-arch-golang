package repositories

import (
	"clean-arch-golang/types"
	"fmt"
)

type MemRepository struct {
	articles map[string]types.Article
}

func NewMemRepository() types.ArticleRepository {
	return &MemRepository{
		articles: make(map[string]types.Article),
	}
}

func (l *MemRepository) Get(title string) (types.Article, error) {
	article, found := l.articles[title]
	if !found {
		return types.Article{}, fmt.Errorf("title not found")
	}
	return article, nil
}

func (l *MemRepository) Articles() ([]types.Article, error) {
	articles := make([]types.Article, 0)
	for _, article := range l.articles {
		articles = append(articles, article)
	}
	return articles, nil
}

func (l *MemRepository) Store(article types.Article) error {
	_, found := l.articles[article.Title]
	if found {
		return fmt.Errorf("title is already exists")
	}
	l.articles[article.Title] = article
	return nil
}

func (l *MemRepository) NumberOfArticles() int {
	return len(l.articles)
}
