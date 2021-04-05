package main

import (
	"fmt"
	"log"
	"os"

	"github.com/sno6/config"
)

type OurConfig struct {
	MyEmail       string            `json:"validate_me" validate:"required,email"`
	SlackWebhooks map[string]string `json:"slack_webhooks" validate:"required"`
}

func main() {
	envVar := os.Getenv("APP_ENV")
	if envVar == "" {
		log.Fatal("Supply an APP_ENV environment variable with a value of `local` or `production`")
	}

	// Pull the environment (local/dev/prod/etc) from the environment.
	env := config.EnvironmentFromString(envVar)

	var cfg OurConfig
	err := config.NewFromFile(&config.Config{
		Path:        "./config",
		Environment: env,
	}, &cfg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("My Email: %v\n", cfg.MyEmail)
	fmt.Printf("Slack Webhooks: %v\n", cfg.SlackWebhooks)
}
