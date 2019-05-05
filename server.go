package main

import (
    "net/http"
    "fmt"
    "log"
    "github.com/pionus/framework"
    "github.com/pionus/framework/middlewares"
)



func main() {
    config := GetConfig()

	app := pionus.NewApp()
    app.Use(middlewares.Logger)
    app.Use(middlewares.Panic)

    app.Static("/static", "./public")

    router := app.Router()

    router.Get("/", func(ctx *pionus.Context) {
		ctx.Text(http.StatusOK, "hey~")
	})

	router.Get(`/hello`, func(ctx *pionus.Context) {
		ctx.Text(http.StatusOK, "Hello world")
	})

	router.Get(`/hello/:name`, func(ctx *pionus.Context) {
        fmt.Printf("hello %s", ctx.Params["name"])
		ctx.Text(http.StatusOK, fmt.Sprintf("Hello %s", ctx.Params["name"]))
	})

    router.Get("/panic", func(ctx *pionus.Context) {
        panic(123)
    })


    err := app.StartTLS(config.Addr, config.Domain...)

	if err != nil {
		log.Fatalf("Could not start server: %s\n", err.Error())
	}

}
