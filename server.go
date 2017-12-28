package main

import (
    "net/http"
    "github.com/pionus/pionus-go/controllers"
)


func main() {
    StaticServ := http.FileServer(http.Dir("public/"))
    MarkdServ := http.FileServer(http.Dir("markdowns/"))
    http.Handle("/assets/", http.StripPrefix("/assets/", StaticServ))
    http.Handle("/md/", http.StripPrefix("/md/", MarkdServ))

    http.HandleFunc("/", controllers.IndexController)
    http.HandleFunc("/articles/", controllers.ArticleController)

    // http.ListenAndServeTLS(":8089", "server.crt", "server.key", nil)
    http.ListenAndServe(":8089", nil)
}
