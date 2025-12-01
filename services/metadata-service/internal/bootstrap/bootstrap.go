package bootstrap

import (
	"context"
	"metadatasvc/internal/config"
	"metadatasvc/internal/db"
	"metadatasvc/internal/repositories"
	"os"
	"shared/pkg/logger"

	"github.com/joho/godotenv"
)

var (
	Config *config.Config
	Logger *logger.Logger
	DB     *db.DB
	Repos  *repositories.Repositories
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

	Logger, err = logger.New(&Config.Log)
	if err != nil {
		return err
	}

	DB, err = db.New(&Config.DB)
	if err != nil {
		return err
	}

	if err = DB.Ping(context.Background()); err != nil {
		return err
	}

	if err = DB.Migrate(); err != nil {
		return err
	}

	Repos = repositories.NewRepositories(DB)

	return nil
}
