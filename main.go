package main

import (
	"fmt"

	"go.uber.org/config"
)

type ApplicationConfig struct {
	Port    int
	Host    string
	BaseUrl string
}

type DatabaseConfig struct {
	Username     string
	Password     string
	Port         int
	Host         string
	DatabaseName string `yaml:"database_name"`
	RequireSsl   string `yaml:"require_ssl"`
}

type EmailClientConfig struct {
	BaseUrl             string `yaml:"base_url"`
	SenderEmail         string `yaml:"sender_email"`
	AuthorizationToken  string `yaml:"authorization_token"`
	TimeoutMilliseconds int    `yaml:"timeout_milliseconds"`
}

type Config struct {
	Database    DatabaseConfig
	Application ApplicationConfig
	EmailClient EmailClientConfig `yaml:"email_client"`
}

func main() {
	opts := []config.YAMLOption{
		config.File("configuration/base.yaml"),
		config.File("configuration/production.yaml"),
	}

	provider, err := config.NewYAML(opts...)
	if err != nil {
		panic(err)
	}

	var c Config
	if err := provider.Get(config.Root).Populate(&c); err != nil {
		panic(err)
	}

	fmt.Printf("config: %v", c)
}
