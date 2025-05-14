package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println()
	fmt.Println(os.Getpid())
	err := os.Mkdir("testdir/anotherdir", 0777)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)

	}

	//Right now only using this sort of stuff for testing.

	//TODO: clear this up later for actual implementation.
	if checkIfFileExists("testdir/anotherdir/test.txt") == false {
		filename := "testdir/anotherdir/test.txt"
		fmt.Printf("%s: file does not exist\n", filename)
	} else {
		fmt.Println("file exists")
	}

}

// In this section we will just declare some basic util functions to aid file and directory management.
func checkIfFileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
