package models

import (
    "time"
)

type Article struct {
    ID string `json:"id"`
    Author string `json:"author"`
    Title string `json:"title"`
    Content string `json:"content"`
    Created time.Time `json:"created"`
}
