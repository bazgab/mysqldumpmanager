package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	fmt.Println()
	fmt.Println(os.Getpid())
	fmt.Println()

	c := []string{"mkdir", "testdir"}
	executeCommand(c)
	fmt.Println("===================")
	loggingInfoForTest := "Info test"
	loggingWarnForTest := "Warn test"
	loggingErrorForTest := "Error test"
	logInfo(loggingInfoForTest)
	logWarn(loggingWarnForTest)
	logError(loggingErrorForTest)
	fmt.Println("===================")
	//Checking if deleteFile function will delete directory
	deleteFile("testdir")
	// Yes it will.

	//Just testing main functionality
	//TODO: This is a very messy way of scripting things, clear up.
	/*fmt.Println("Checking for testfile.txt...")
	if checkIfFileExists("testfile.txt") == false {
		fmt.Println("Testfile.txt does not exist, creating...")
		//	createFile("testfile.txt")
		filename := "testfile.txt"
		content := "This is the text content for testfile.txt\n" +
			"With multiple lines\n"

		createFileWithContent(filename, content)
	}
	fmt.Println("Checking again for Testfile.txt...")
	if checkIfFileExists("testfile.txt") == true {
		fmt.Println("Found testfile.txt")
	}*/

}

// TODO: We need to write every logging action to a log file, not just print out the error in the console

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
	logInfo(r.String())
	err := r.Run()
	if err != nil {
		m := ": Command could not be executed, this can happen if the directory is already created on your system"
		logError(err.Error() + m)
	}
	fmt.Println("Executed command: mkdir testdir ")
}
