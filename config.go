package main

import "github.com/caarlos0/env"

// Config holds configuration values from the environment, with sane defaults
// (where possible). Required configuration will throw a Fatal error if they
// are missing.
type Config struct {
	AppName string `env:"APP_NAME" envDefault:"gl2pp"`
	AppEnv  string `env:"APP_ENV" envDefault:"development"`
	AppDesc string `env:"APP_DESC" envDefault:"gitlab to planningpoker.com exporter"`
}

// NewConfig returns an instance of Config, with values loaded from ENV vars.
func NewConfig() (Config, error) {
	c := Config{}
	err := env.Parse(&c)
	return c, err
}
