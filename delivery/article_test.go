package delivery

import (
	"bytes"
	"clean-arch-golang/mocks"
	"clean-arch-golang/types"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func Test_ArticleHandler(t *testing.T) {
	wroteAt := time.Now()
	mockArticle := types.Article{
		ID:        1,
		Title:     "title0",
		Content:   "content0",
		CreatedAt: wroteAt,
	}
	mockArticleUsecase := new(mocks.ArticleUsecase)
	mockArticleUsecase.On("Store", mockArticle).Return(nil)
	mockArticleUsecase.On("NumberOfArticles").Return(1)

	var testRequest = func(req *http.Request) *httptest.ResponseRecorder {
		r := gin.Default()
		NewArticleHandler(r, mockArticleUsecase)
		rec := httptest.NewRecorder()
		r.ServeHTTP(rec, req)
		return rec
	}

	t.Run("should get article using title", func(t *testing.T) {
		mockArticleUsecase.On("Get", "title0").Return(mockArticle, nil)
		req, err := http.NewRequest(http.MethodGet, "/article/title0", strings.NewReader(""))
		require.Nil(t, err)
		resp := testRequest(req)
		require.Equal(t, 200, resp.Code)
	})

	t.Run("should reject invalid title", func(t *testing.T) {
		mockArticleUsecase.On("Get", "title1").Return(types.Article{}, fmt.Errorf("invalid title"))
		req, err := http.NewRequest(http.MethodGet, "/article/title1", strings.NewReader(""))
		require.Nil(t, err)
		resp := testRequest(req)
		require.Equal(t, 400, resp.Code)
	})

	t.Run("should store article", func(t *testing.T) {
		mockArticleUsecase.On("Store", mock.AnythingOfType("types.Article")).Return(nil)
		bin, err := json.Marshal(mockArticle)
		require.Nil(t, err)
		req, err := http.NewRequest(http.MethodPost, "/article", bytes.NewBuffer(bin))
		require.Nil(t, err)
		resp := testRequest(req)
		require.Equal(t, 200, resp.Code)
	})

	t.Run("should get articles", func(t *testing.T) {
		mockArticleUsecase.On("Articles").Return([]types.Article{mockArticle}, nil)
		req, err := http.NewRequest(http.MethodGet, "/articles", strings.NewReader(""))
		require.Nil(t, err)
		resp := testRequest(req)
		require.Equal(t, 200, resp.Code)
	})
}
