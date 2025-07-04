package cmd

import (
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates mysqldumpmanager based on the configuration file",
	Long: `Updates mysqldumpmanager based on the configuration file. 

By default, the configuration file is located at /etc/mysqldumpmanager/conf.yaml`,
	Run: updateCmdFunc,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateCmdFunc(cmd *cobra.Command, args []string) {

	LogInfo("Looking for conf.yaml and updating MySQLDumpManager")

	type Config struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	}

	var configValues Config

	yamlFile, err := os.ReadFile("/etc/mysqldumpmanager/conf.yaml")
	if err != nil {
		LogError(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &configValues)
	if err != nil {
		LogError(err.Error())
	}

	LogInfo("Fetching data from conf.yaml...")

	err = os.Setenv("MYSQLDUMPMANAGER_USER", configValues.User)
	if err != nil {
		LogError(err.Error())
	}
	LogInfo("Setting up environment variable MYSQLDUMPMANAGER_USER")

	err = os.Setenv("MYSQLDUMPMANAGER_PASSWORD", configValues.Password)
	if err != nil {
		LogError(err.Error())
	}
	LogInfo("Setting up environment variable MYSQLDUMPMANAGER_PASSWORD")

	UserVariableString := os.Getenv("MYSQLDUMPMANAGER_USER")
	PasswordVariableString := os.Getenv("MYSQLDUMPMANAGER_PASSWORD")

	// Converting it to a String slice
	uByte := []byte(UserVariableString)
	pByte := []byte(PasswordVariableString)

	m := []byte("\n[mariadb-dump]\nuser= ")
	m2 := []byte("\npassword= ")
	m3 := []byte("\n")
	// Now we have to append in groups and then finally append it all together
	a1 := append(m, uByte...)
	a2 := append(m2, pByte...)
	a3 := append(a1, a2...)
	fByteAppend := append(a3, m3...)
	f, err := os.Create("/etc/my.cnf.d/mariadb-dump.cnf")
	if err != nil {
		LogError(err.Error())
	}
	_, err = f.Write(fByteAppend)
	if err != nil {
		LogError(err.Error())
	}
	LogInfo("Created/Updated /etc/my.cnf.d/mariadb-dump.cnf")
	err = f.Close()
	if err != nil {
		LogError(err.Error())
	}
}
