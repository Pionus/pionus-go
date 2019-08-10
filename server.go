package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/pionus/arry"
    "github.com/pionus/arry/middlewares"

    "github.com/pionus/pionus-go/controllers"
    "github.com/pionus/pionus-go/graphql"
)



func main() {
	app := arry.New()
    app.Use(middlewares.Gzip)
    app.Use(middlewares.LoggerToFile("logs/access.log"))
    app.Use(middlewares.Panic)
    app.Use(middlewares.Auth(Config.Authorization))

    app.Static("/md", "./markdowns")
    app.Static("/assets", "./theme/"+ Config.Theme +"/assets")
    app.Static("/node_modules", "./theme/"+ Config.Theme +"/node_modules")
    app.Static("/web_modules", "./theme/"+ Config.Theme +"/web_modules")
    app.Views("./theme/"+ Config.Theme +"/pages")

    router := app.Router()

    router.Get("/", controllers.IndexController)
    router.Get("/article", controllers.ArticleList)
    router.Get("/article/:id", controllers.ArticleDetail)

    router.Post("/wp-json/wp/v2/posts", controllers.WPPost)

	router.Get(`/hello`, func(ctx arry.Context) {
		ctx.Text(http.StatusOK, "Hello world")
	})

	router.Get(`/hello/:name`, func(ctx arry.Context) {
        fmt.Printf("hello %s", ctx.Param("name"))
		ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s", ctx.Param("name")))
	})

    router.Get("/push", func(ctx arry.Context) {
        ctx.Push("/assets/styles/main.css")
        ctx.Push("/assets/article.js")
        ctx.Text(http.StatusOK, "pushed~")
    })

    router.Post("/graphql", graphql.GetController())


    err := app.Start(Config.Addr)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
