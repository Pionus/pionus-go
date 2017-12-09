package controllers

import (
    "net/http"
    "html/template"
)

// import (
//     "github.com/kataras/iris/mvc"
// )
//
// type IndexController struct {
//     mvc.C
// }
//
// func (c *IndexController) Get() mvc.Result {
//     return mvc.View {
//         Name: "index.html",
//         Data: "",
//     }
// }

func IndexController (w http.ResponseWriter, r *http.Request) {
    t, _ := template.ParseFiles("views/index.html")
    err := t.Execute(w, struct{}{})

    if err != nil {

    }
}
