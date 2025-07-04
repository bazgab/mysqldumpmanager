/*

	This command is meant to be run only once, during first startup of the program. It does the following things
	in order:
		1. Check for the right directory/configuration file
		2. If it's not present, creates the missing directories/files

	An improvement we could make is to have a different output if the "mysqldumpmanager init" command has already
	been executed before, instead of outputting the check like a first time startup.

*/

// TODO: Change the 'mysql' string literals to 'mariadb'

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

	LogInfo("Checking for MySQLDumpManager directory...")
	p := "/etc/mysqldumpmanager/"
	if _, err := os.Stat(p); os.IsNotExist(err) {
		LogWarn("/etc/mysqldumpmanager/ directory does not exist")
		LogInfo("Attempting to create /etc/mysqldumpmanager/ directory...")
		CreateDirectory(p)
	} else {
		LogInfo("/etc/mysqldumpmanager/ directory exists")
	}

	LogInfo("Checking for configuration file...")
	f := "/etc/mysqldumpmanager/conf.yaml"
	if CheckIfFileExists(f) == false {
		LogError("Configuration file does not exist")
		LogWarn("/etc/mysqldumpmanager/config.yaml is required, attempting to create it")
		LogInfo("Creating /etc/mysqldumpmanager/config.yaml...")
		createConfigurationFile()

	}
	LogInfo("Configuration file exists")

	LogInfo("Checking for /etc/mysqldumpmanager/dumps directory...")
	d := "/etc/mysqldumpmanager/dumps"
	if _, err := os.Stat(d); os.IsNotExist(err) {
		LogWarn("/etc/mysqldumpmanager/dumps directory does not exist")
		LogInfo("Attempting to create /etc/mysqldumpmanager/dumps directory...")
		CreateDirectory(d)
	} else {
		LogInfo("/etc/mysqldumpmanager/dumps directory exists")
	}

	LogInfo("Success - Init check completed. Run 'mysqldumpmanager --help' for usage.")
}

func CreateDirectory(filepath string) {
	LogInfo("Creating directory " + filepath)
	err := os.MkdirAll(filepath, 0755)
	if err != nil {
		LogError(err.Error())
	}

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
