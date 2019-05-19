package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/pionus/arry"
    "github.com/pionus/arry/middlewares"
)



func main() {
    config := GetConfig()

	app := arry.New()
    app.Use(middlewares.Logger)
    app.Use(middlewares.Panic)

    app.Static("/static", "./public")

    router := app.Router()

    router.Get("/", func(ctx arry.Context) {
		ctx.Text(http.StatusOK, "hey~")
	})

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


    err := app.Start(config.Addr)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
