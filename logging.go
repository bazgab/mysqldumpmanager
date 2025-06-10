package main

import (
	"bytes"
	"fmt"
	"log"
)

// Log INFO Block
func logInfo(info string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "[INFO] ", log.Ltime|log.Lmsgprefix)
	)

	err := logger.Output(2, info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}

// Log WARN Block
func logWarn(warn string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "[WARN] ", log.Ltime|log.Lmsgprefix)
	)

	err := logger.Output(2, warn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}

// Log ERROR Block
func logError(error string) {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "[ERROR] ", log.Ltime|log.Lmsgprefix)
	)

	err := logger.Output(2, error)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}
