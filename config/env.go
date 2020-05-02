package config

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spf13/afero"
)

type EnvOptions struct {
	// TODO: use util path
	ConfigPath string
}

type Env struct {
	// Init while construct
	Fs            afero.Fs
	BackgroundCtx context.Context

	// Init in InitEnvFromOptions
	ConfigPath string
	Cfg        Config
}

func InitEnvFromOptions(env *Env, options *EnvOptions) error {
	env.ConfigPath = options.ConfigPath

	cfg, err := read(env.ConfigPath)
	if err != nil {
		return err
	}
	env.Cfg = cfg

	return nil
}

func read(configPath string) (Config, error) {
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		return Config{}, err
	}

	configJSON, err := ioutil.ReadAll(configFile)
	if err != nil {
		return Config{}, err
	}

	var config Config
	if err := json.Unmarshal(configJSON, &config); err != nil {
		return Config{}, err
	}

	return config, nil
}
