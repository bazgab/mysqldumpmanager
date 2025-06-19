package internal

import (
	"fmt"
	"github.com/bazgab/mysqldumpmanager/cmd"
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// TODO: properly implement config file creation on init.go file
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

	n := configValues.Name
	fmt.Println("Name:", n)
	fmt.Printf("yaml test values file property Age: %v \n", configValues.Age)

	n1 := os.Getenv("MYSQL_USER")
	fmt.Println("Using environment variable MYSQL_USER with value: ", n1)
	err = os.Setenv("TEST_SET_ENV", configValues.Name)
	if err != nil {
		return
	}
	fmt.Println("Setting up environment variable TEST_SET_ENV using yamltest.yaml, value: ", os.Getenv("TEST_SET_ENV"))

}

// TODO: Create a method for checking if config files are being properly handled.
func checkConfigFile() {
	// This function will check if the config file is present, for testing we need to create something on the tests pkg.

}
