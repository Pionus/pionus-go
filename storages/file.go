package main

import (
    "os"
    "time"
    "regexp"
    "strings"
    "io/ioutil"
    "path/filepath"
    "github.com/pionus/pionus-go/models"
)

type service struct {
    basePath string
}

var Service = service{
    basePath: "markdowns/",
}

func (s service) GetArticleByID(id string) (*models.Article, error) {
    file, err := ioutil.ReadFile("markdowns/" + id + ".md")
    if err != nil {
        return nil, err
    }

    title := parseTitle(file)
    t, _ := time.Parse("20060102150405", id)

    return &models.Article{
        ID: id,
        Author: "Secbone",
        Title: title,
        Content: string(file),
        Created: t,
    }, nil
}

func (s service) GetArticleList() (*[]*models.Article, error) {
    var list []*models.Article

    err := filepath.Walk(s.basePath, func(path string, info os.FileInfo, err error) error {
        if info.IsDir() {
            return nil
        }

        fullname := info.Name()
        ext := filepath.Ext(fullname)

        if ext == ".md" {
            name := strings.TrimSuffix(fullname, ext)
            article, _ := s.GetArticleByID(name)
            list = append(list, article)
        }

        return nil
    })

    if err != nil {
        panic(err)
    }

    return &list, nil
}

func (s service) SaveArticle(a *models.Article) (*models.Article, error) {
    now := time.Now()
    m := models.Article{
        ID: now.Format("20060102150405"),
        Content: "# " + a.Title + "\n" + a.Content,
    }

    f, err := os.OpenFile(s.basePath + m.ID + ".md", os.O_CREATE | os.O_EXCL | os.O_WRONLY, 0644)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    f.WriteString(m.Content)

    return &m, nil
}


func parseTitle(file []byte) string {
    re := regexp.MustCompile(`^#\s+(.+)\s`)
    title := re.FindSubmatch(file)
    return string(title[1])
}

