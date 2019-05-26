package controllers

import (
    "github.com/pionus/arry"
)

type Article struct {
    Id  string
}

func ArticleController(ctx arry.Context) {
    ctx.Push("/assets/marked.min.js")
    ctx.Push("/assets/article.js")
    ctx.Push("/md/20171209.md")

    article :=  Article{
        Id: ctx.Param("id"),
    }
    ctx.Render(200, "article.html", article)
}
