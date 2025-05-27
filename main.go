package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println()
	fmt.Println(os.Getpid())
	/*
		err := os.Mkdir("testdir/anotherdir", 0777)
		if err != nil && !os.IsExist(err) {
			log.Fatal(err)

		}
	*/

	//Just testing main functionality
	//TODO: This is a very messy way of scripting things, clear up.
	fmt.Println("Checking for testfile.txt...")
	if checkIfFileExists("testfile.txt") == false {
		fmt.Println("Testfile.txt does not exist, creating...")
		createFile("testfile.txt")
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
// This is gonna be used for creating a /etc/mysqldumpmanager file for storing .sql dumps
func createFile(filename string) {
	_, err := os.Create(filename)
	if err != nil {
		return
	}

}

//Delete file:
// Used for deleting dump files

func deleteFile(filename string) {
	err := os.Remove(filename)
	if err != nil {
		return
	} else {
		fmt.Println("Deleted testfile.txt")
	}
}
