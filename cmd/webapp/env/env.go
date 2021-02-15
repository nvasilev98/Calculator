package env

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
}

func Load() (Config, error) {
	var cfg Config
	if err := ReadEnvVariable("HOST", &cfg.Host); err != nil {
		return Config{}, err
	}
	if err := ReadEnvVariable("PORT", &cfg.Port); err != nil {
		return Config{}, err
	}
	if err := ReadEnvVariable("USERNAME", &cfg.Username); err != nil {
		return Config{}, err
	}
	if err := ReadEnvVariable("PASSWORD", &cfg.Password); err != nil {
		return Config{}, err
	}
	return cfg, nil
}
