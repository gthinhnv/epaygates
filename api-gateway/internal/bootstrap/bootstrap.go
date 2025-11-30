package bootstrap

import (
	"apigateway/internal/config"
	"os"

	"github.com/joho/godotenv"
)

var (
	Config *config.Config
)

func Init() error {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		return err
	}

	Config, err = config.Load(os.Getenv("ENV"))
	if err != nil {
		return err
	}

	return nil
}
