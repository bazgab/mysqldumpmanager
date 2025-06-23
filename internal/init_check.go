package internal

import (
	"github.com/bazgab/mysqldumpmanager/cmd"
	"os"
)

func InitCheckEntryPoint() {
	CheckForMySQLDumpManagerDirectory()
	CheckForConfigurationFile()
}

func CheckForMySQLDumpManagerDirectory() {
	cmd.LogInfo("Checking for MySQLDumpManager directory...")
	p := "/etc/mysqldumpmanager/"
	if _, err := os.Stat(p); os.IsNotExist(err) {
		cmd.LogWarn("/etc/mysqldumpmanager/ directory does not exist")
		cmd.LogInfo("Attempting to create /etc/mysqldumpmanager/ directory...")
		CreateMySQLDumpManagerDirectory()
	} else {
		cmd.LogInfo("/etc/mysqldumpmanager/ directory exists")
	}
}

func CheckForConfigurationFile() {
	cmd.LogInfo("Checking for configuration file...")
	f := "/etc/mysqldumpmanager/config.yaml"
	cmd.CheckIfFileExists(f)
}
