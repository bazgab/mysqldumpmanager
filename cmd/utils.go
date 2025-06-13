package cmd

import (
	"os"
	"os/exec"
	"strings"
)

// In this section we will just declare some basic util functions to aid file and directory management.
// Check if file exists
func checkIfFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

// Create File:
// This is going to be used for creating a /etc/mysqldumpmanager file for storing .sql dumps
// and to create the file for storing MySQL credentials so we don't need to prompt password everytime
func createFileWithContent(filename string, content string) {
	f, err := os.Create(filename)
	if err != nil {
		logError(err.Error())
	}
	_, err = f.WriteString(content)
	if err != nil {
		logError(err.Error())
	}
}

// Delete file:
// Used for deleting dump files

func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		m := ": File could not be deleted, check if file exists"
		logError(err.Error() + m)
	} else {
		m2 := "File successfully deleted"
		logInfo(m2)
	}
}

// Execute command
// This will be used to execute the mysqldump command to create new dumps
func executeCommand(cmd []string) {

	r := exec.Command(cmd[0], cmd[1:]...)
	err := r.Run()
	if err != nil {
		m := ": Command could not be executed, this can happen if the directory is already created on your system"
		logError(err.Error() + m)
	}
	cmdStr := strings.Join(cmd, " ")
	m := "Executed command successfully: " + "\"" + cmdStr + "\""
	logInfo(r.String())
	logInfo(m)
}
