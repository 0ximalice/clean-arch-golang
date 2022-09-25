package delivery

import (
	"clean-arch-golang/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	articleUsecase types.ArticleUsecase
}

func NewArticleHandler(r *gin.Engine, us types.ArticleUsecase) {
	h := articleHandler{
		articleUsecase: us,
	}
	r.GET("/article/:title", h.FetchArticle)
	r.GET("/articles", h.FetchArticles)
	r.POST("/article", h.StoreArticle)
}

func (h *articleHandler) FetchArticle(c *gin.Context) {
	article, err := h.articleUsecase.Get(c.Param("title"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "NOK",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "OK",
		"article": article,
	})
}

func (h *articleHandler) FetchArticles(c *gin.Context) {
	articles, err := h.articleUsecase.Articles()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "NOK",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":   "OK",
		"articles": articles,
	})
}

func (h *articleHandler) StoreArticle(c *gin.Context) {
	var article types.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "NOK",
			"error":  err.Error(),
		})
		return
	}
	err = h.articleUsecase.Store(article)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "NOK",
			"error":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "OK",
	})
}
