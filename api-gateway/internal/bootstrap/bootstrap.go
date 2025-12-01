package bootstrap

import (
	"apigateway/internal/config"
	"os"
	sharedConfig "shared/config"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	SharedConfig *sharedConfig.Config
	Config       *config.Config

	MetadataServiceConn *grpc.ClientConn
)

func Init() error {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		return err
	}

	appEnv := os.Getenv("ENV")

	SharedConfig, err = sharedConfig.Load(appEnv)
	if err != nil {
		return err
	}

	Config, err = config.Load(appEnv)
	if err != nil {
		return err
	}

	MetadataServiceConn, err = grpc.NewClient(SharedConfig.MetadataService.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	return nil
}
