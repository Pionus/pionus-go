package main

import (
    "log"
    "os"
    "strings"
    "path/filepath"
)

func main() {
    var files []string

    root := "markdowns/"

    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

        if info.IsDir() {
            return nil
        }

        fullname := info.Name()
        ext := filepath.Ext(fullname)
        if ext == ".md" {
            name := strings.TrimSuffix(fullname, ext)
            files = append(files, name)
        }

        return nil
    })

    if err != nil {
        panic(err)
    }
}
