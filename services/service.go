package services

import (
    "log"
    "plugin"
    "github.com/pionus/pionus-go/models"
)

type service interface {
    GetArticleByID(id string) (*models.Article, error)
    GetArticleList() (*[]*models.Article, error)
    SaveArticle(a *models.Article) (*models.Article, error)
}

func getService() service {
    plug, err := plugin.Open("./storages/file.so")

    if err != nil {
        panic(err)
    }

    s, err := plug.Lookup("Service")
    if err != nil {
        panic(err)
    }

    ser, ok := s.(service)
    if !ok {
        log.Print("ssssssssss")
    }

    return ser
}

var Service = getService()
