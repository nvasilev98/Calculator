package connection

import "github.com/nvasilev98/calculator/cmd/webapp/env"

type ConfigDatabase struct {
	Host     string
	Port     string
	Username string
	Password string
	Name     string
}

func Load() (ConfigDatabase, error) {
	var cfg ConfigDatabase
	if err := env.ReadEnvVariable("DB_HOST", &cfg.Host); err != nil {
		return ConfigDatabase{}, err
	}
	if err := env.ReadEnvVariable("DB_PORT", &cfg.Port); err != nil {
		return ConfigDatabase{}, err
	}
	if err := env.ReadEnvVariable("DB_USERNAME", &cfg.Username); err != nil {
		return ConfigDatabase{}, err
	}
	if err := env.ReadEnvVariable("DB_PASSWORD", &cfg.Password); err != nil {
		return ConfigDatabase{}, err
	}
	if err := env.ReadEnvVariable("DB_NAME", &cfg.Name); err != nil {
		return ConfigDatabase{}, err
	}
	return cfg, nil
}
