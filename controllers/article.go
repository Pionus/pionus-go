package controllers

import (
    "github.com/pionus/arry"
    "github.com/pionus/pionus-go/services"
)

func ArticleList(ctx arry.Context) {
    ctx.Render(200, "index.html", nil)
}

func ArticleDetail(ctx arry.Context) {
    id := ctx.Param("id")
    article, _ := services.GetArticleByID(id)

    ctx.Render(200, "article.html", article)
}
