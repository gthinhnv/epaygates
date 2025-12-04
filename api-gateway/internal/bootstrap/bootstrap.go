package bootstrap

import (
	"apigateway/internal/config"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	sharedConfig "shared/config"
	"shared/pkg/logger"
	"shared/pkg/translator"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	SharedConfig *sharedConfig.Config
	Config       *config.Config

	Logger *logger.Logger

	Translator *translator.Translator

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

	Translator = translator.New(buildLocaleEntries(Config.SupportedLangs), Config.DefaultLang)

	MetadataServiceConn, err = grpc.NewClient(SharedConfig.MetadataService.GRPCAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}

	return nil
}

func buildLocaleEntries(supportedLangs []string) []*translator.LocaleEntry {
	_, filename, _, _ := runtime.Caller(0)
	baseDir := filepath.Dir(filename)

	localeEntries := make([]*translator.LocaleEntry, len(supportedLangs))

	for i, lang := range supportedLangs {
		localeEntries[i] = &translator.LocaleEntry{
			Name: lang,
			Path: filepath.Join(baseDir, fmt.Sprintf("../locales/%s.json", lang)),
		}
	}

	return localeEntries
}
