package config

import (
    "io/ioutil"
    "encoding/json"
    "log"
)

type GTConfig struct {
    Id  string  `json:"id"`
    Key string  `json:"key"`
}

type RootConfig struct {
    Geetest GTConfig `json:"geetest"`
}

func LoadConfig(path string) (RootConfig, error) {
    var c RootConfig
    raw, err := ioutil.ReadFile(path)

    if err != nil {
        log.Fatal(err)
        return c, err
    }

    json.Unmarshal(raw, &c)
    return c, nil
}

var Config, _ = LoadConfig("config.json")
