package main

import (
	"flag"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"github.com/joho/godotenv"
	"github.com/keys4words/go-restfull_api/standardWebserver/internal/app/api"
)

var (
	configPath string
	format     string
)

func init() {
	flag.StringVar(&format, "format", ".toml", "format of config file")
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	flag.Parse()
	log.Println("It works")
	config := api.NewConfig()

	if format == ".toml" {
		_, err := toml.DecodeFile(configPath, config)
		if err != nil {
			log.Println("cannot parse toml file, use default value ", err)
		}
	}
	if format == ".env" {
		err := godotenv.Load(configPath)
		if err != nil {
			log.Println("cannot parse .env file, use default value ", err)
		}
		config.BindAddr = os.Getenv("bind_addr")
		config.LoggerLevel = os.Getenv("logger_level")
	}

	server := api.New(config)
	log.Fatal(server.Start())
}
