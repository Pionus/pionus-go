package controllers

import (
    // "fmt"
    // "io/ioutil"
    // "net/http"
    // "html/template"

    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
)

type ArticleController struct {
    Ctx iris.Context
}

type Article struct {
    Id  string
}

func (c *ArticleController) Get() mvc.Result {

    // options := &http.PushOptions{
    //     Header: http.Header{},
    // }

    c.Ctx.ResponseWriter().Push("/assets/marked.min.js", nil)
    c.Ctx.ResponseWriter().Push("/assets/article.js", nil)

    return mvc.View {
        Name: "article.html",
        Data: Article{},
    }
}

func (c *ArticleController) GetBy(id string) mvc.Result {

    c.Ctx.ResponseWriter().Push("/assets/marked.min.js", nil)
    c.Ctx.ResponseWriter().Push("/assets/article.js", nil)

    return mvc.View {
        Name: "article.html",
        Data: Article{
            Id: id,
        },
    }
}

// func ArticleController(w http.ResponseWriter, r *http.Request) {
//     pusher, ok := w.(http.Pusher)
//
//     if ok {
//         fmt.Printf("Push is supported")
//         options := &http.PushOptions{
//             Header: http.Header{
//                 "Accept-Encoding": r.Header["Accept-Encoding"],
//             },
//         }
//
//         pusher.Push("/assets/marked.min.js", options)
//         pusher.Push("/assets/article.js", options)
//         pusher.Push("/md/20171209.md", options)
//
//     }
//
//     // file, err := ioutil.ReadFile("markdowns/20171209.md")
//     // body := template.HTML(blackfriday.Run(file))
//
//     t, _ := template.ParseFiles("views/article.html")
//     err := t.Execute(w, Article{})
//
//     if err != nil {
//
//     }
// }
