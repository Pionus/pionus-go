package main

import (
    // "fmt"
    // "context"
    // "io/ioutil"
    // "encoding/json"
    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"

    "github.com/pionus/pionus-go/controllers"
    "github.com/pionus/pionus-go/resolver"
    "github.com/pionus/pionus-go/schema"
    "github.com/pionus/pionus-go/middlewares/graphql"

    // graphql "github.com/graph-gophers/graphql-go"
    // "github.com/graph-gophers/graphql-go/relay"
    // "log"
)

func main() {
    app := iris.New()

    graph := graphql.New(graphql.Options{
        Schema: schema.String(),
        Resolver: &resolver.Resolver{},
    })

    // pugEngine := iris.Pug("./views", ".pug")
    amberEngine := iris.Amber("./views", ".amber")
    htmlEngine := iris.HTML("./views", ".html")
    amberEngine.Reload(true)
    app.RegisterView(amberEngine)
    app.RegisterView(htmlEngine)

    // crtpath := "server.crt"
    // keypath := "server.key"

    app.StaticWeb("/assets/", "./public/")
    app.StaticWeb("/md/", "./markdowns/")

    mvc.New(app.Party("/")).Handle(new(controllers.IndexController))
    mvc.New(app.Party("/articles/")).Handle(new(controllers.ArticleController))
    mvc.New(app.Party("/dreams/")).Handle(new(controllers.DreamsController))
    mvc.New(app.Party("/editor/")).Handle(new(controllers.EditorController))

    app.Any("/graphql/", graph)

    app.Run(
        iris.Addr(":8089"),
        // iris.TLS(":8089", crtpath, keypath),
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
