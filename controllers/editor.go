package controllers

import (
    // "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
)

type EditorController struct {}

func (c *EditorController) Get() mvc.Result {
    return mvc.View {
        Name: "editor.html",
    }
}
