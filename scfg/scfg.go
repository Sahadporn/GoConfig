package scfg

import (
	"github.com/caarlos0/env/v6"
)

type Config struct {
	Env    string `env:"ENV_VAR_STRUCT" envDefault:"HI"`
	Number int    `env:"MAGIC_NUMBER" envDefault:"72"`
}

// New returns a new Config struct
func New() Config {
	cfg := Config{}
	env.Parse(&cfg)
	return cfg
}
