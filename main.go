package main

import (
    "github.com/kataras/iris"
    "github.com/pionus/pionus-go/controllers"
)

func main() {
    app := iris.New()

    app.StaticWeb("/assets", "./public")
    app.RegisterView(iris.HTML("./views", ".html"))
    app.Controller("/", new(controllers.IndexController))

    app.Run(iris.Addr(":8099"))
}
