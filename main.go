package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {

	fmt.Println()
	fmt.Println(os.Getpid())
	fmt.Println()
	// TODO: Fix this

	executeCommand()
	fmt.Println("===================")

	//Just testing main functionality
	//TODO: This is a very messy way of scripting things, clear up.
	fmt.Println("Checking for testfile.txt...")
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
	}

}

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
		log.Fatal(err)
	}
	_, err = f.WriteString(content)
	if err != nil {
		log.Fatal(err)
	}
}

// Delete file:
// Used for deleting dump files

func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		return
	} else {
		fmt.Println("Deleted testfile.txt")
	}
}

// Execute command
// This will be used to execute the mysqldump command to create new dumps
func executeCommand() {
	cmd := exec.Command("mkdir", "testdir")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Executed command: mkdir testdir ")
}
