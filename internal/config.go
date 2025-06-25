package internal

import (
	"github.com/bazgab/mysqldumpmanager/cmd"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func configSetup() {
	var configValues Config

	yamlFile, err := os.ReadFile("/etc/mysqldumpmanager/conf.yaml")
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &configValues)
	if err != nil {
		cmd.LogError(err.Error())
	}

	u := configValues.User
	p := configValues.Password
	cmd.LogInfo("Fetching data from conf.yaml...")
	cmd.LogInfo("User: " + u)
	cmd.LogInfo("Password: " + p)

	err = os.Setenv("MYSQLDUMPMANAGER_USER", configValues.User)
	if err != nil {
		cmd.LogError(err.Error())
	}
	cmd.LogInfo("Setting up environment variable MYSQLDUMPMANAGER_USER with value: " + os.Getenv("MYSQLDUMPMANAGER_USER"))

	err = os.Setenv("MYSQLDUMPMANAGER_PASSWORD", configValues.Password)
	if err != nil {
		cmd.LogError(err.Error())
	}
	cmd.LogInfo("Setting up environment variable MYSQLDUMPMANAGER_PASSWORD with value: " + os.Getenv("MYSQLDUMPMANAGER_PASSWORD"))

}
