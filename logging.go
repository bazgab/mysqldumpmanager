package main

import (
	"bytes"
	"fmt"
	"log"
)

// TODO: Fix this
func loginfo(info string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "INFO: ", log.Lshortfile)
	)

	err := logger.Output(2, info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}

// using this for testing
func main() {
	info := "Hello World!"
	loginfo(info)
}
