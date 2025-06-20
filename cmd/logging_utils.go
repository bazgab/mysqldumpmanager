package cmd

import (
	"bytes"
	"fmt"
	"github.com/fatih/color"
	"log"
)

func LogInfo(info string) {

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

func LogWarn(warn string) {

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

func LogError(error string) {

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
