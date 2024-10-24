package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	APIserver "github.com/Sunsh1ne0/rest-api-golang.git/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to apiserver configuration")
}

func main() {
	flag.Parse()

	config := APIserver.NewConfig()

	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatalf("Error decoding configuration: %v", err)
	}

	s := APIserver.New(config)

	if err := s.Start(); err != nil {
		log.Fatalf("Error starting API server: %v", err)
	}
}
