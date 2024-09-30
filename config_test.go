package main

import (
	"fmt"
	"github.com/golang-template/pkg/types"
	"gopkg.in/yaml.v2"
	"os"
	"testing"
)

type AddTest struct {
	description string
	configSet   config
	configError string
	configName  string
}

var addTests = []AddTest{
	{
		description: "Valid Config with output to logfile.",
		configSet: config{
			Log: logS{
				Level: "debug",
				Out:   types.Fileout,
				Name:  "test_log_01.log",
			},
		},
		configError: "",
		configName:  "test_config01.yml",
	},
	{
		description: "Not valid config. Log File Name is empty.",
		configSet: config{
			Log: logS{
				Level: "debug",
				Out:   types.Fileout,
				Name:  "",
			},
		},
		configError: "ERROR reading config. Config doesn't have log.name or it is empty!",
		configName:  "test_config02.yml",
	},
	{
		description: "Valid config. To stdout.",
		configSet: config{
			Log: logS{
				Level: "debug",
				Out:   types.Stdout,
				Name:  "",
			},
		},
		configError: "",
		configName:  "test_config03.yml",
	},
	{
		description: "Valid config. To stdout. level info",
		configSet: config{
			Log: logS{
				Level: "info",
				Out:   types.Stdout,
				Name:  "",
			},
		},
		configError: "",
		configName:  "test_config04.yml",
	},
	{
		description: "Valid config. To stdout. level error",
		configSet: config{
			Log: logS{
				Level: "error",
				Out:   types.Stdout,
				Name:  "",
			},
		},
		configError: "",
		configName:  "test_config05.yml",
	},
	{
		description: "Valid config. To stdout. level error",
		configSet: config{
			Log: logS{
				Level: "asfasfsf",
				Out:   types.Stdout,
				Name:  "",
			},
		},
		configError: "",
		configName:  "test_config06.yml",
	},
	{
		description: "Not Valid config. Name is not set at all.",
		configSet: config{
			Log: logS{
				Level: "debug",
				Out:   types.Stdout,
			},
		},
		configError: "",
		configName:  "test_config22.yml",
	},
}

func TestConfigInit(t *testing.T) {
	fmt.Println(len(addTests))
	for _, test := range addTests {
		wantConfig, err := createConfig(test.configName, test.configSet)
		if err != nil {
			t.Errorf("Config not created\n")
		}

		getConfig, err := configInit(test.configName)
		if err != nil {
			if err.Error() != test.configError {
				t.Errorf("not expceted error. Expect [%s] but received [%s]\n",
					test.configError, err.Error())
			}
		} else {
			if *wantConfig != *getConfig {
				t.Errorf("Config parsing Error. want %v get %v\n", wantConfig, *getConfig)
			}
		}
	}

}

func createConfig(filename string, cfg config) (*config, error) {

	// Marshal the struct to YAML format
	data, err := yaml.Marshal(&cfg)
	if err != nil {
		fmt.Printf("Error while marshaling to YAML: %v\n", err)
		return nil, err
	}

	// Write to the c.yml file
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		fmt.Printf("Error while writing to file: %v\n", err)
		return nil, err
	}

	fmt.Printf("Successfully written to %s\n\n", filename)
	return &cfg, nil
}
