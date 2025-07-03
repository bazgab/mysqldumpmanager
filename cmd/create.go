package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a dump",
	Long: `Creates a dump. 

Dumps are located in /etc/mysqldumpmanager/dumps/`,
	Run: createCmdFunc,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("file", "F", "", "file name to write dump to (required)")
	err := createCmd.MarkFlagRequired("file")
	if err != nil {
		LogError(err.Error())
	}
}

func createCmdFunc(cmd *cobra.Command, args []string) {

	// Change directory
	dFilePath := "/etc/mysqldumpmanager/dumps/"
	err := os.Chdir(dFilePath)
	if err != nil {
		LogError("Couldn't change directory: " + err.Error())
	}

	// Validate
	wDir, _ := os.Getwd()
	if err != nil {
		LogError("Couldn't change working directory: " + err.Error())
	}
	LogInfo("Setting Working directory as: " + wDir)

	// TODO: Fix functionality on the mysqldump executable (/usr/bin/mysql)
	dCmd := []string{"/usr/bin/mysqldump", "test_db", ">", "test-dump1.sql"}
	cmdStr := strings.Join(dCmd, " ")

	executeCommand(cmdStr)
	LogInfo("Executed command: " + cmdStr)

}
