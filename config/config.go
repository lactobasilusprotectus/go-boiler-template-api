package config

import (
	"fmt"
	"github.com/joeshaw/envdecode"
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	Http HttpConfig
}

//HttpConfig untuk config server
type HttpConfig struct {
	Port    string `env:"APP_PORT"`
	Timeout int    `env:"TIMEOUT"`
}

type DatabaseConfig struct {
	Driver            string `env:"DB_CONNECTION"`
	Host              string `env:"DB_HOST"`
	User              string `env:"DB_USERNAME"`
	Password          string `env:"DB_PASSWORD"`
	Database          string `env:"DB_DATABASE"`
	Port              string `env:"DB_PORT"`
	MaxOpenConnection int    `env:"DB_MAX_OPEN_CONNECTION"`
	MaxIdleConnection int    `env:"DB_MAX_IDLE_CONNECTION"`
}

//Read => fungsi untuk membaca file config
func Read() (cfg Config, err error) {
	if err = godotenv.Load(); err != nil {
		log.Println("err reading .env")
	}

	if err = envdecode.StrictDecode(&cfg); err != nil {
		err = fmt.Errorf("erro decoding config: %+v", err)
		return
	}

	return cfg, nil
}
