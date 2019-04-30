package main

type Config struct {
    Addr string
    Cert string
    Key string
}

func GetConfig() *Config {
    return &Config{
        Addr: ":80",
        Cert: "server.crt",
        Key: "server.key",
    }
}
