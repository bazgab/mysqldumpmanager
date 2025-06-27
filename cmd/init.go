package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes MySQLDumpManager",
	Long: `Initializes MySQLDumpManager. This command is meant to be used only once after installation, 
and provides the instructions for MySQLDumpManager to create the necessary directories and files to make
it work as intended.`,
	Run: initFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initFunc(cmd *cobra.Command, args []string) {

}

func CheckForMySQLDumpManagerDirectory() {
	LogInfo("Checking for MySQLDumpManager directory...")
	p := "/etc/mysqldumpmanager/"
	if _, err := os.Stat(p); os.IsNotExist(err) {
		LogWarn("/etc/mysqldumpmanager/ directory does not exist")
		LogInfo("Attempting to create /etc/mysqldumpmanager/ directory...")
		CreateMySQLDumpManagerDirectory()
	} else {
		LogInfo("/etc/mysqldumpmanager/ directory exists")
	}
}

func CheckForConfigurationFile() {
	LogInfo("Checking for configuration file...")
	f := "/etc/mysqldumpmanager/conf.yaml"
	if CheckIfFileExists(f) == false {
		LogError("Configuration file does not exist")
		LogWarn("/etc/mysqldumpmanager/config.yaml is required, attempting to create it")
		LogInfo("Creating /etc/mysqldumpmanager/config.yaml...")
		createConfigurationFile()

	}
	LogInfo("Configuration file exists")
}

func CreateMySQLDumpManagerDirectory() {
	LogInfo("Creating directory for MySQLDumpManager at /etc/mysqldumpmanager...")
	err := os.MkdirAll("/etc/mysqldumpmanager", 0755)
	if err != nil {
		LogError(err.Error())
	}

	// First time checking if directory was created successfully
	if _, err := os.Stat("/etc/mysqldumpmanager"); err != nil {
		if os.IsNotExist(err) {
			LogError("Failure - Unable to create directory...")
			LogError(err.Error())
		}

	} else {
		LogInfo("Success - /etc/mysqldumpmanager directory exists")
	}
}

func createConfigurationFile() {
	LogInfo("Creating configuration file...")
	f := "/etc/mysqldumpmanager/conf.yaml"
	c := []byte("#Authentication\nuser:\npassword:\n")
	CreateFileWithContent(f, c)
}
