package models

import (
    "time"
)

type Dream struct {
    Id int
    Author string
    Face string
    Content string
    Created time.Time
    Year int
    Month time.Month
    Day int
}
