package configs

import (
	"fmt"
	"strings"

	env "github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Secrets struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
	Db       string `env:"POSTGRES_DB"`
	Ssl      string `env:"POSTGRES_SSL_MODE"`
}

func GetSecret() (secret Secrets) {

	// Load env from local
	if err := godotenv.Load(strings.Join([]string{"./configs", ".env"}, "/")); err != nil {
		panic(fmt.Errorf("unable to load .env file: %v", err))
	}

	if err := env.Parse(&secret); err != nil {
		panic(fmt.Errorf("unable to parse environment variables: %v", err))
	}

	return secret
}
