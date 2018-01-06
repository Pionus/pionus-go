package main

import (
    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
    "github.com/pionus/pionus-go/controllers"
)


func main() {
    app := iris.New()

    app.RegisterView(iris.HTML("./views", ".html"))
    crtpath := "server.crt"
    keypath := "server.key"

    app.StaticWeb("/assets/", "./public/")
    app.StaticWeb("/md/", "./markdowns/")

    mvc.New(app.Party("/")).Handle(new(controllers.IndexController))
    mvc.New(app.Party("/articles/")).Handle(new(controllers.ArticleController))

    app.Run(
        iris.TLS(":8089", crtpath, keypath),
    )

    // StaticServ := http.FileServer(http.Dir("public/"))
    // MarkdServ := http.FileServer(http.Dir("markdowns/"))
    // http.Handle("/assets/", http.StripPrefix("/assets/", StaticServ))
    // http.Handle("/md/", http.StripPrefix("/md/", MarkdServ))
    //
    // http.HandleFunc("/", controllers.IndexController)
    // http.HandleFunc("/articles/", controllers.ArticleController)
    //
    // http.ListenAndServeTLS(":8089", crtpath, keypath, nil)
    // http.ListenAndServe(":8089", nil)
}
