package controllers

import (
    // "fmt"
    "time"
    "database/sql"
    "crypto/md5"
    "encoding/hex"
    "github.com/kataras/iris"
    "github.com/kataras/iris/mvc"
    _ "github.com/go-sql-driver/mysql"
    // _ "github.com/mattn/go-sqlite3"
    "github.com/pionus/pionus-go/models"
)

type DreamsController struct {
    Ctx iris.Context
}

type List struct {
    Dreams []models.Dream
}

type Response struct {
    Code int `json:"code"`
    Message string `json:"message"`
    Data interface{} `json:"data"`
}

type DreamData struct {
    Name string `json:"name"`
}


func (c *DreamsController) Get() mvc.Result {
    db, _ := getDB()

    defer db.Close()

    rows, _ := db.Query("select * from dreams")

    var dreams []models.Dream

    for rows.Next() {
        d, _ := readDream(rows)

        dreams = append(dreams, d)
    }

    return mvc.View {
        Name: "dreams.amber",
        Data: List{
            Dreams: dreams,
        },
    }
}

func (c *DreamsController) GetTest() {
    c.Ctx.JSON(Response{
        Code: 200,
        Message: "ok",
        Data: DreamData{
            Name: "123",
        },
    })
}

func (c *DreamsController) Post() {
    content := c.Ctx.PostValue("content")
    email := c.Ctx.PostValue("email")
    db, _ := getDB()
    defer db.Close()

    stmt, _ := db.Prepare("insert into dreams (content, author) values (?, ?)")
    stmt.Exec(content, email)

    c.Ctx.JSON(Response{
        Code: 200,
        Message: "ok",
        Data: nil,
    })
}

func getDB() (*sql.DB, error) {
    // return sql.Open("sqlite3", "sqlite/dreams.db")
    return sql.Open("mysql", "root@/dreams?parseTime=true")
}


func readDream(rows *sql.Rows) (models.Dream, error) {
    var (
        id int
        content string
        author string
        created time.Time
    )

    err := rows.Scan(&id, &author, &content, &created)

    return models.Dream{
        Id: id,
        Author: author,
        Face: getHash(author),
        Content: content,
        Created: created,
        Year: created.Year(),
        Month: created.Month(),
        Day: created.Day(),
    }, err
}

func getHash(source string) string {
    hasher := md5.New()
    hasher.Write([]byte(source))
    return hex.EncodeToString(hasher.Sum(nil))
}
