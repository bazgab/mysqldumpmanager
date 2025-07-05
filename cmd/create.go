package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"os/exec"
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
	createCmd.Flags().StringP("database", "D", "", "database to dump (required)")
	// Marking as required
	err := createCmd.MarkFlagRequired("file")
	if err != nil {
		LogError(err.Error())
	}

	err = createCmd.MarkFlagRequired("database")
	if err != nil {
		LogError(err.Error())
	}
}

func createCmdFunc(cmd *cobra.Command, args []string) {

	// TODO: Test if it works on /etc/
	// Change directory
	dFilePath := os.Getenv("HOME") + "/mysqldumpmanager/test-dumps"
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

	cmdStr := "mariadb-dump mariadb_test > test-dump1.sql"
	LogInfo("Executing: " + cmdStr)
	c := exec.Command("/bin/sh", "-c", "mariadb-dump mariadb_test > test-dump1.sql")

	var stdout []byte
	stdout, err = c.Output()

	if err != nil {
		LogError(err.Error())
	}

	// Print the output
	LogInfo(string(stdout))

}
