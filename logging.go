package main

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"log"
)

// Log INFO Block
func logInfo(info string) {

	c := color.New(color.FgBlue, color.Bold).SprintFunc()
	p := c("[INFO] ")

	var (
		buf    bytes.Buffer
		logger = log.New(&buf, p, log.Ltime|log.Lmsgprefix)
	)

	err := logger.Output(2, info)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}

// Log WARN Block
func logWarn(warn string) {

	c := color.New(color.FgYellow, color.Bold).SprintFunc()
	p := c("[WARN] ")

	var (
		buf    bytes.Buffer
		logger = log.New(&buf, p, log.Ltime|log.Lmsgprefix)
	)

	err := logger.Output(2, warn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}

// Log ERROR Block
func logError(error string) {

	c := color.New(color.FgRed, color.Bold).SprintFunc()
	p := c("[ERROR] ")

	var (
		buf    bytes.Buffer
		logger = log.New(&buf, p, log.Ltime|log.Lmsgprefix)
	)

	err := logger.Output(2, error)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(&buf)
}
