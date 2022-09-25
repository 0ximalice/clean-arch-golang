package main

import (
	"clean-arch-golang/delivery"
	"clean-arch-golang/repositories"
	"clean-arch-golang/usecases"

	"github.com/gin-gonic/gin"
)

func main() {
	articleRepository := repositories.NewMemRepository()
	articleUsecase := usecases.NewArticleUsecase(articleRepository)

	r := gin.Default()
	delivery.NewArticleHandler(r, articleUsecase)
	r.Run()
}
