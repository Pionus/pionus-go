package services

import (
    "os"
    "time"
    "regexp"
    "strings"
    "io/ioutil"
    "path/filepath"
    "github.com/pionus/pionus-go/models"
)


func parseTitle(file []byte) string {
    re := regexp.MustCompile(`^#\s+(.+)\s`)
    title := re.FindSubmatch(file)
    return string(title[1])
}


func GetArticleByID(id string) (*models.Article, error) {
    file, err := ioutil.ReadFile("markdowns/" + id + ".md")
    title := parseTitle(file)

    if err != nil {
        return nil, err
    }

    return &models.Article{
        ID: id,
        Author: "Secbone",
        Title: title,
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
