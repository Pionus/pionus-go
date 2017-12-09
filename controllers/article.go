package controllers

import (
    "io/ioutil"
    "net/http"
    "html/template"

    "gopkg.in/russross/blackfriday.v2"
)

type Article struct {
    Body    template.HTML
}

func ArticleController(w http.ResponseWriter, r *http.Request) {
    file, err := ioutil.ReadFile("markdowns/20171209.md")
    body := template.HTML(blackfriday.Run(file))

    t, _ := template.ParseFiles("views/article.html")
    err = t.Execute(w, Article{ body })

    if err != nil {

    }
}
