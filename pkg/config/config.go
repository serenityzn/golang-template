package config

import (
	"github.com/golang-template/pkg/types"
	"github.com/spf13/viper"
)

const portMin = 1024
const portMax = 49151

type logS struct {
	Level types.LogLevel  `yaml:"level"`
	Out   types.LogOutput `yaml:"out"`
	Name  types.LogName   `yaml:"name"`
}

type http struct {
	Host    types.HttpHost     `yaml:"host"`
	Port    types.ConfHttpPort `yaml:"port"`
	Timeout int                `yaml:"timeout"`
}

type config struct {
	Log  *logS `yaml:"log"`
	Http *http `yaml:"http"`
}

func (c *config) Validate() bool {
	if c.Http == nil || c.Log == nil || c.Http.Timeout == 0 || c.Http.Host.Validate() != nil || c.Http.Port > portMax ||
		c.Http.Port < portMin || c.Log.Out > 1 || c.Log.Out < 0 || len(c.Log.Name) == 0 {
		return false
	}
	return true
}
func ConfigInit(configName string) (*config, error) {

	viper.SetConfigName(configName)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		return nil, err
	}

	conf := &config{}
	err = viper.Unmarshal(conf)
	if err != nil {
		return nil, err
	}

	return conf, nil
}
