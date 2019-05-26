package models

import (
    "time"
)

type Article struct {
    ID string
    Author string
    Title string
    Content string
    Created time.Time
}
