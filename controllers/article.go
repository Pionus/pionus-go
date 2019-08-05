package controllers

import (
    "github.com/pionus/arry"
)

type Article struct {
    Id  string
}

func ArticleList(ctx arry.Context) {
    ctx.Render(200, "article.html", nil)
}

func ArticleDetail(ctx arry.Context) {
    article :=  Article{
        Id: ctx.Param("id"),
    }
    ctx.Render(200, "article.html", article)
}
