package main

import (
    "io/ioutil"
    "encoding/json"
)

type Config struct {
    Addr string `json:"addr"`
    Cert string `json:"cert"`
    Key string `json:"key"`
}

func GetConfig() *Config {
    file, _ := ioutil.ReadFile("config.json")
    var config Config
    json.Unmarshal(file, &config)
    return &config
}
