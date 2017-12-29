package controllers

import (
    "fmt"
    // "io/ioutil"
    "net/http"
    "html/template"

    // "gopkg.in/russross/blackfriday.v2"
)

type Article struct {
    Body    template.HTML
}

func ArticleController(w http.ResponseWriter, r *http.Request) {
    pusher, ok := w.(http.Pusher)

    if ok {
        fmt.Printf("Push is supported")
        options := &http.PushOptions{
            Header: http.Header{
                "Accept-Encoding": r.Header["Accept-Encoding"],
            },
        }

        pusher.Push("/assets/marked.min.js", options)
        pusher.Push("/assets/article.js", options)
        pusher.Push("/md/20171209.md", options)

    }

    // file, err := ioutil.ReadFile("markdowns/20171209.md")
    // body := template.HTML(blackfriday.Run(file))

    t, _ := template.ParseFiles("views/article.html")
    err := t.Execute(w, Article{})

    if err != nil {

    }
}
