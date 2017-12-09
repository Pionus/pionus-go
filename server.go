package main

import (
    "net/http"
    "github.com/pionus/pionus-go/controllers"
)


func main() {
    StaticServ := http.FileServer(http.Dir("public/"))
    http.Handle("/assets/", http.StripPrefix("/assets/", StaticServ))

    http.HandleFunc("/", controllers.IndexController)
    http.HandleFunc("/articles/", controllers.ArticleController)

    http.ListenAndServe(":8089", nil)
}
