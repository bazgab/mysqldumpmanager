package internal

import (
	"github.com/bazgab/mysqldumpmanager/cmd"
	"os"
)

func initEnvSetup() {

	initConfigMessage := "Initializing configuration for MySQLDumpManager"
	cmd.LogInfo(initConfigMessage)

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
	c := []byte("#Authentication\nuser: root\npassword: rootpasswd\n")
	cmd.CreateFileWithContent(f, c)
}
