package config

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type app struct {
	Host  string `required:"true" yaml:"host"`
	Port  int    `required:"true" yaml:"port"`
	Debug bool   `required:"true" yaml:"debug"`
}

type database struct {
	Dialect  string `default:"postgres" yaml:"dialect"`
	Database string `required:"true" yaml:"database"`
	Debug    bool   `default:"false" yaml:"debug"`
	Username string `required:"true" yaml:"username"`
	Password string `required:"true" yaml:"password"`
	Host     string `required:"true" yaml:"host"`
	Port     int    `required:"true" yaml:"port"`
	SSLMode  bool   `default:"false" yaml:"ssl_mode"`
}

type configuration struct {
	App      app      `yaml:"app"`
	Database database `yaml:"database"`
}

var App app
var Database database

func InitConfiguration() {
	var configName string
	// Look for a flag called config to set the proper config file.
	// Uses a default value of `production`.
	flag.StringVar(&configName, "config", "production", "This flag defines which file should to be taken")
	flag.Parse()

	var configFileName string
	switch configName {
	case "production":
		configFileName = "production.config.yaml"
	case "dev":
		configFileName = "dev.config.yaml"
	default:
		// Stop the app if the flag was set incorrectly.
		log.Fatalf("Not valid environment flag, use `dev` or `production` instead.\n")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error getting the current word directory: \n%v\n", err)
	}

	configFile, err := os.Open(fmt.Sprintf("%s/config/%s", wd, configFileName))
	if err != nil {
		log.Fatalf("The config file %s has not been possible to open: \n%v\n", configFileName, err)
	}

	// Parse the content inside an configuration struct.
	var configuration configuration
	d := yaml.NewDecoder(configFile)
	err = d.Decode(&configuration)
	if err != nil {
		log.Fatalf("Error while decoding config file into structure, be sure the structure of the file is ok: \n%v\n", err)
	}

	// Set two exportable objects based on the config gotten before.
	App = configuration.App
	Database = configuration.Database

	log.Printf("Configuration setted\n")
}
