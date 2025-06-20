package internal

import (
	"github.com/bazgab/mysqldumpmanager/cmd"
	"os"
)

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
