package services

import (
    "os"
    "time"
    "strings"
    "io/ioutil"
    "path/filepath"
    "github.com/pionus/pionus-go/models"
)

func GetArticleByID(id string) (*models.Article, error) {
    file, err := ioutil.ReadFile("markdowns/" + id + ".md")

    if err != nil {
        return nil, err
    }

    return &models.Article{
        ID: id,
        Author: "Secbone",
        Title: id,
        Content: string(file),
        Created: time.Now(),
    }, nil
}

func GetArticleList() (*[]*models.Article, error) {
    var list []*models.Article

    err := filepath.Walk("markdowns/", func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }

        fullname := info.Name()
        ext := filepath.Ext(fullname)

        if ext == ".md" {
            name := strings.TrimSuffix(fullname, ext)
            article, _ := GetArticleByID(name)
            list = append(list, article)
        }

        return nil
    })

    if err != nil {
        panic(err)
    }

    return &list, nil
}
