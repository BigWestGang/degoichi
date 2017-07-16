package main

import (
    "fmt"
    "github.com/BurntSushi/toml"
    "log"
)

type Config struct {
  Provider string
  Box string
}

type Configs struct {
  Config []Config
}

func main() {
  var config Configs
  _, err := toml.DecodeFile("config.toml", &config)
  if err != nil {
    log.Fatal(err)
  }
  for _, s := range config.Config {
    fmt.Printf("%s (%s)\n", s.Provider, s.Box)
  }
}
