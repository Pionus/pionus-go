package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/pionus/arry"
    "github.com/pionus/arry/middlewares"

    "./controllers"
    "./graphql"
)



func main() {
    config := GetConfig()

	app := arry.New()
    app.Use(middlewares.Logger)
    app.Use(middlewares.Panic)

    app.Static("/assets", "./public")
    app.Static("/md", "./markdowns")
    app.Views("views/")

    router := app.Router()

    router.Get("/", controllers.IndexController)
    router.Get("/article/:id", controllers.ArticleController)

	router.Get(`/hello`, func(ctx arry.Context) {
		ctx.Text(http.StatusOK, "Hello world")
	})

	router.Get(`/hello/:name`, func(ctx arry.Context) {
        fmt.Printf("hello %s", ctx.Param("name"))
		ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s", ctx.Param("name")))
	})

    router.Get("/push", func(ctx arry.Context) {
        ctx.Push("/static/article.css")
        ctx.Push("/static/article.js")
        ctx.Text(http.StatusOK, "pushed~")
    })

    router.Post("/graphql", graphql.GetController())


    err := app.Start(config.Addr)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
