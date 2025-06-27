package internal

/*
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
	f := "/etc/mysqldumpmanager/conf.yaml"
	if cmd.CheckIfFileExists(f) == false {
		cmd.LogError("Configuration file does not exist")
		cmd.LogWarn("/etc/mysqldumpmanager/config.yaml is required, attempting to create it")
		cmd.LogInfo("Creating /etc/mysqldumpmanager/config.yaml...")
		createConfigurationFile()

	}
	cmd.LogInfo("Configuration file exists")
}

func CreateMySQLDumpManagerDirectory() {
	cmd.LogInfo("Creating directory for MySQLDumpManager at /etc/mysqldumpmanager...")
	err := os.MkdirAll("/etc/mysqldumpmanager", 0755)
	if err != nil {
		cmd.LogError(err.Error())
	}

	// First time checking if directory was created successfully
	if _, err := os.Stat("/etc/mysqldumpmanager"); err != nil {
		if os.IsNotExist(err) {
			cmd.LogError("Failure - Unable to create directory...")
			cmd.LogError(err.Error())
		}

	} else {
		cmd.LogInfo("Success - /etc/mysqldumpmanager directory exists")
	}
}

func createConfigurationFile() {
	cmd.LogInfo("Creating configuration file...")
	f := "/etc/mysqldumpmanager/conf.yaml"
	c := []byte("#Authentication\nuser:\npassword:\n")
	cmd.CreateFileWithContent(f, c)
}

*/
