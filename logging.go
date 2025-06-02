package main

import (
	"bytes"
	"fmt"
	"log"
)

// Log INFO Block
func loginfo(info string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "[INFO] ", log.Ldate|log.Ltime|log.Llongfile)
	)

	err := logger.Output(2, info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}

// Log WARN Block
func logwarn(warn string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "[WARN] ", log.Ldate|log.Ltime|log.Llongfile)
	)

	err := logger.Output(2, warn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}

// Log ERROR Block
func logerror(error string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "[ERROR] ", log.Ldate|log.Ltime|log.Llongfile)
	)

	err := logger.Output(2, error)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}
