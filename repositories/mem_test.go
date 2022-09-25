package repositories

import (
	"clean-arch-golang/types"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func Test_MemRepository(t *testing.T) {
	memRepository := NewMemRepository()
	wroteAt := time.Now()
	testArticle := types.Article{
		ID:        1,
		Title:     "title0",
		Content:   "content0",
		CreatedAt: wroteAt,
	}

	t.Run("should store", func(t *testing.T) {
		err := memRepository.Store(testArticle)
		require.Nil(t, err)
	})

	t.Run("should list", func(t *testing.T) {
		articles, err := memRepository.Articles()
		require.Nil(t, err)
		require.Equal(t, 1, len(articles))
	})

	t.Run("should fetch number of articles", func(t *testing.T) {
		require.Equal(t, 1, memRepository.NumberOfArticles())
	})

	t.Run("should get using title", func(t *testing.T) {
		article, err := memRepository.Get(testArticle.Title)
		require.Nil(t, err)
		require.EqualValues(t, testArticle, article)
	})
}
