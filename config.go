package main

import (
    "io/ioutil"
    "encoding/json"
)

type config struct {
    Addr string `json:"addr"`
    Cert string `json:"cert"`
    Key string `json:"key"`
    Storage string `json:"storage"`
    Authorization string `json:"authorization"`
    Theme string `json:"theme"`
}

func GetConfig() *config {
    file, _ := ioutil.ReadFile("config.json")
    var c config
    json.Unmarshal(file, &c)
    return &c
}


var Config = GetConfig()
