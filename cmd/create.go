package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
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
	/*
		if err != nil {
			c := color.New(color.FgRed, color.Bold).SprintFunc()
			p := c("[ERROR] ")
			fmt.Println(p, "Command failed to execute")
		}
	*/

	// Change directory

	// Change the working directory

	err := os.Chdir("/etc/mysqldumpmanager/dumps")
	if err != nil {
		LogError("Couldn't change directories" + err.Error())
	}

	// Verify the new working directory
	wd, _ := os.Getwd()
	if err != nil {
		LogError(err.Error())
	}
	fmt.Println("Working directory: " + wd)

	// Execute mysqldump command

}
