package main

import (
	"fmt"
	"log"

	"github.com/sno6/config"
)

type OurConfig struct {
	FromJSON string `json:"from_json" validate:"required"`
	FromEnv  string `env:"FROM_ENV" validate:"required"`
}

func main() {
	var cfg OurConfig
	err := config.NewFromFile(&config.Config{
		Path:        "./config",
		Environment: config.Local,
	}, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("FromJSON: %v\n", cfg.FromJSON)
	fmt.Printf("FromEnv: %v\n", cfg.FromEnv)
}
