package cmd

import (
	"fmt"
)

func main() {

	fmt.Println("ENTRY POINT: ")

	c := []string{"mkdir", "testdir"}
	executeCommand(c)
	fmt.Println("===================")

	// Insert sample messages for testing
	loggingInfoForTest := "Info test"
	loggingWarnForTest := "Warn test"
	loggingErrorForTest := "Error test"

	logInfo(loggingInfoForTest)
	logWarn(loggingWarnForTest)
	logError(loggingErrorForTest)
	fmt.Println("===================")

	//Checking if deleteFile function will delete directory
	fmt.Println("Deleting testdir to not hog tests")
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
