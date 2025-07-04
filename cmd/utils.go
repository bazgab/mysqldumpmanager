package cmd

import (
	"os"
	"os/exec"
)

func CheckIfFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func CreateFileWithContent(filename string, content []byte) {
	err := os.WriteFile(filename, content, 0755)
	if err != nil {
		LogError(err.Error())
	}
}

func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		m := ": File could not be deleted, check if file exists"
		LogError(err.Error() + m)
	} else {
		m2 := "File successfully deleted"
		LogInfo(m2)
	}
}

func executeCommand(cmd string) {

	r := exec.Command(cmd)
	err := r.Run()
	if err != nil {
		m := ": Command could not be executed"
		LogError(err.Error() + m)
	}
	m := "Executed command successfully: " + "\"" + cmd + "\""
	LogInfo(r.String())
	LogInfo(m)
}
