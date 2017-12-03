package controllers

import (
    "github.com/kataras/iris/mvc"
)

type IndexController struct {
    mvc.C
}

func (c *IndexController) Get() mvc.Result {
    return mvc.View {
        Name: "index.html",
        Data: "",
    }
}
