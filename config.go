package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/pborman/getopt/v2"
	"gopkg.in/yaml.v2"
)

type CmdArgs struct {
	ConfigFile string
	Args       []string
}

type AppConfig struct {
	Name string `yaml:"name"`
}

func parseArgs() CmdArgs {
	helpFlag := getopt.BoolLong("help", 'h', "显示帮助")
	configFlag := getopt.StringLong("config", 'c', "config.yml", "指定配置文件")
	getopt.Parse()
	args := getopt.Args()

	if *helpFlag {
		getopt.Usage()
		os.Exit(0)
	}

	result := CmdArgs{
		ConfigFile: *configFlag,
		Args:       args,
	}
	return result
}

func loadConfig(args CmdArgs) AppConfig {
	configFile, err := ioutil.ReadFile(args.ConfigFile)
	if err != nil {
		log.Fatalln("加载配置文件失败", err)
	}

	var appConfig AppConfig
	yaml.Unmarshal(configFile, &appConfig)

	return appConfig
}
