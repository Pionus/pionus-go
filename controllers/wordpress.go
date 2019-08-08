package controllers

import (
    "github.com/pionus/arry"
    "github.com/pionus/pionus-go/models"
    "github.com/pionus/pionus-go/services"
)

type PostResponse struct {
    Id string `json:"id"`
    Link string `json:"link"`
}


func WPPost(ctx arry.Context) {
    if !ctx.Get("auth").(bool) {
        ctx.Reply(403)
        return
    }

    var p models.Article
    ctx.Decode(&p)
    err := services.SaveArticle(&p)

    if err != nil {
        ctx.Reply(500)
        return
    }

    ctx.JSON(200, PostResponse{
        Id: "123",
        Link: "localhost",
    })
}
