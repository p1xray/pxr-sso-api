package app

import (
	"flag"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"os"
	grpcclient "pxr-sso-api/internal/client/grpc"
)

const (
	commonConfigFlagName    = "config"
	commonConfigFlagUsage   = "path to config file"
	defaultCommonConfigPath = "./config/config.yaml"

	environmentConfigFlagName  = "env-config"
	environmentConfigFlagUsage = "path to environment config file"

	envVarCommonConfigPath      = "PXR_SSO_API_CONFIG_PATH"
	envVarEnvironmentConfigPath = "PXR_SSO_API_ENV_CONFIG_PATH"
)

// Config is the application configuration.
type Config struct {
	Env         string            `yaml:"env" env:"PXR_SSO_API_ENV" env-default:"local" env-upd:""`
	Port        string            `yaml:"port" evn:"PXR_SSO_API_PORT" env-required:"true" env-upd:""`
	GRPCClients grpcclient.Config `yaml:"grpc" env-required:"true"`
}

type configLoader struct {
	cfg *Config

	commonConfigPath      string
	environmentConfigPath string
}

func newConfigLoader() *configLoader {
	return &configLoader{
		cfg: &Config{},
	}
}

// MustLoad loads config and panics if any error occurs.
func (l *configLoader) MustLoad() *Config {
	l.fetchPath()

	if err := l.loadCommonConfig(); err != nil {
		panic(err)
	}

	_ = l.loadEnvironmentConfig()

	return l.cfg
}

// fetchPath fetches path from command line flag or environment variable.
// Priority: flag > env > default.
func (l *configLoader) fetchPath() {
	l.flagParse()
	l.envVarParse()
}

func (l *configLoader) flagParse() {
	var commonConfigPath string
	var environmentConfigPath string
	flag.StringVar(&commonConfigPath, commonConfigFlagName, defaultCommonConfigPath, commonConfigFlagUsage)
	flag.StringVar(&environmentConfigPath, environmentConfigFlagName, "", environmentConfigFlagUsage)
	flag.Parse()

	l.setCommonConfigPathIfEmpty(commonConfigPath)
	l.setEnvironmentConfigPathIfEmpty(environmentConfigPath)
}

func (l *configLoader) envVarParse() {
	_ = godotenv.Load()

	commonConfigPath := os.Getenv(envVarCommonConfigPath)
	l.setCommonConfigPathIfEmpty(commonConfigPath)

	environmentConfigPath := os.Getenv(envVarEnvironmentConfigPath)
	l.setEnvironmentConfigPathIfEmpty(environmentConfigPath)
}

func (l *configLoader) setCommonConfigPathIfEmpty(path string) {
	if l.commonConfigPath == "" {
		l.commonConfigPath = path
	}
}

func (l *configLoader) setEnvironmentConfigPathIfEmpty(path string) {
	if l.environmentConfigPath == "" {
		l.environmentConfigPath = path
	}
}

func (l *configLoader) loadCommonConfig() error {
	if l.commonConfigPath == "" {
		return fmt.Errorf("common config path is empty")
	}

	if err := l.loadByPath(l.commonConfigPath); err != nil {
		return err
	}

	return nil
}

func (l *configLoader) loadEnvironmentConfig() error {
	if l.environmentConfigPath == "" {
		return fmt.Errorf("environment config path is empty")
	}

	if err := l.loadByPath(l.environmentConfigPath); err != nil {
		return err
	}

	return nil
}

func (l *configLoader) loadByPath(path string) error {
	file, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("%s: %s", "config file does not exists", path)
	}

	if err = cleanenv.ReadConfig(path, l.cfg); err != nil {
		return fmt.Errorf("%s %s: %w", "cannot read", file.Name(), err)
	}

	return nil
}
