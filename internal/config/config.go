package config

import "github.com/joho/godotenv"

type GRPCConfig interface {
	Address() string
}

type PgConfig interface {
	DSN() string
}

type HTTPConfig interface {
	Address() string
}

type SwaggerConfig interface {
	Address() string
}

func Load(path string) error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
