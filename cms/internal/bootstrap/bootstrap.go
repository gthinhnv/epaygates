package bootstrap

import (
	"apigateway/pkg/utils/modelutil"
	"cms/internal/config"
	"os"
	sharedConfig "shared/config"
	"shared/models/commonmodel"
	"shared/pkg/logger"

	"github.com/joho/godotenv"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	SharedConfig *sharedConfig.Config
	Config       *config.Config

	Logger *logger.Logger

	JSON jsoniter.API

	PageTypeItems []*commonmodel.PageTypeItem
	StatusItems   []*commonmodel.StatusItem

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

	Logger, err = logger.New(&Config.Log)
	if err != nil {
		return err
	}

	JSON = jsoniter.ConfigCompatibleWithStandardLibrary

	StatusItems = modelutil.BuildStatuses()
	PageTypeItems = modelutil.BuildPageTypes()

	MetadataServiceConn, err = grpc.NewClient(SharedConfig.MetadataService.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	return nil
}
