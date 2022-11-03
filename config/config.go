package config

import (
	"errors"
	"path"
	"runtime"
	"strings"

	"github.com/iamolegga/enviper"
	"github.com/rest-api/logger"
	"github.com/spf13/viper"
)

// we want to use the current working directory as the config directory
func getPwd() string {
	_, b, _, _ := runtime.Caller(0)
	return path.Join(path.Dir(path.Dir(path.Dir(b))))
}

// GetConfig configuration
func GetConfig() (Config, error) {
	var cfg Config

	e := enviper.New(viper.New())

	e.AddConfigPath(getPwd())
	e.SetConfigName(".env")
	e.SetConfigType("env")

	// enable viper to handle env values for nested structs
	e.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// defaults to ENV variable values
	e.AutomaticEnv()

	if err := e.ReadInConfig(); err != nil {
		logger.Warn(err)
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			logger.Fatal("Error reading config file: ", err)
		}
	}

	err := e.Unmarshal(&cfg)
	if err != nil {
		logger.Errorf("error to decode, %v", err)
		return cfg, nil
	}

	// just a bare minimum check, port will be provided either through
	// the env file
	// or by export SERVER_PORT=8000
	if cfg.Server.Port == "" {
		return cfg, errors.New("error reading config")
	}

	return cfg, nil
}
