package controllers

import (
    "github.com/pionus/arry"
)


func IndexController (ctx arry.Context) {
    ctx.Render(200, "blog.html", nil)
}
