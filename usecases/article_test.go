package usecases

import (
	"clean-arch-golang/repositories"
	"clean-arch-golang/types"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_ArticleUsecase(t *testing.T) {
	mockRepo := repositories.NewMemRepository()
	articleUsecase := NewArticleUsecase(mockRepo)

	testArticle := types.Article{
		Title:   "title0",
		Content: "content0",
	}

	t.Run("should store", func(t *testing.T) {
		err := articleUsecase.Store(testArticle)
		require.Nil(t, err)
	})

	t.Run("should list", func(t *testing.T) {
		articles, err := articleUsecase.Articles()
		require.Nil(t, err)
		require.Equal(t, 1, len(articles))
	})

	t.Run("should get using title", func(t *testing.T) {
		article, err := articleUsecase.Get(testArticle.Title)
		require.Nil(t, err)

		t.Run("should fill an id and updated_at", func(t *testing.T) {
			require.NotEqual(t, 0, article.ID)
			require.NotEqual(t, "", article.CreatedAt)
		})
	})
}
