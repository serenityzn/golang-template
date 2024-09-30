package main

import (
	"errors"
	"github.com/golang-template/pkg/types"
	"github.com/spf13/viper"
)

type logS struct {
	Level string          `yaml:"level"`
	Out   types.LogOutput `yaml:"out"`
	Name  string          `yaml:"name"`
}

type http struct {
	HostAddress string `yaml:"hostaddress"`
	Timeout     int    `yaml:"timeout"`
}

type config struct {
	Log  logS `yaml:"log"`
	Http http `yaml:"http"`
}

func configInit(configName string) (*config, error) {
	var myConfig = config{}

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	logOut := viper.GetInt("log.out")
	if logOut != 0 && logOut != 1 {
		return nil, errors.New("ERROR reading config. Config doesn't have log.out value!")
	}

	if logOut == int(types.Stdout) {
		myConfig.Log.Out = types.Stdout
	} else if logOut == int(types.Fileout) {
		myConfig.Log.Out = types.Fileout
		fileName := viper.GetString("log.name")
		if len(fileName) < 1 {
			return nil, errors.New("ERROR reading config. Config doesn't have log.name or it is empty!")
		}
		myConfig.Log.Name = fileName
	}

	logLevel := viper.GetString("log.level")
	if len(logLevel) < 1 {
		return nil, errors.New("ERROR reading config. Config doesn't have log.level value!!!")
	}

	myConfig.Log.Level = "info"
	if logLevel == "info" || logLevel == "debug" || logLevel == "error" {
		myConfig.Log.Level = logLevel
	}

	myConfig.Http.HostAddress = viper.GetString("http.hostaddress")
	myConfig.Http.Timeout = viper.GetInt("http.timeout")

	return &myConfig, nil
}
