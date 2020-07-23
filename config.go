package main

import (
	"log"

	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	// Title string
	DB database `toml:"database"`
}

type database struct {
	Server        string
	Blog_db       string
	Blog_text_col string
}

var Config tomlConfig

func init() {
	if _, err := toml.DecodeFile("config.toml", &Config); err != nil {
		log.Fatal(err)
	}
	log.Println(Config)

}
