package controllers

import (
    "github.com/pionus/arry"

    // "gopkg.in/russross/blackfriday.v2"
)

type Article struct {
    Id  string
}

func ArticleController(ctx arry.Context) {
    ctx.Push("/assets/marked.min.js")
    ctx.Push("/assets/article.js")
    ctx.Push("/md/20171209.md")

    article :=  Article{
        Id: "20171209",
    }
    ctx.Render(200, "article.html", article)


    // file, err := ioutil.ReadFile("markdowns/20171209.md")
    // body := template.HTML(blackfriday.Run(file))
}
