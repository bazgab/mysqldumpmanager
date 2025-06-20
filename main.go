package main

import (
	"github.com/bazgab/mysqldumpmanager/cmd"
	"github.com/bazgab/mysqldumpmanager/internal"
)

func main() {
	internal.CheckForMySQLDumpManagerDirectory()
	cmd.Execute()
}
