package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a dump",
	Long: `Creates a dump. 

By default, dumps are located in /etc/mysqldumpmanager/dumps/`,
	Run: createCmdFunc,
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().StringP("file", "F", "sample.sql", "file to be created")
}

func createCmdFunc(cmd *cobra.Command, args []string) {
	/*
		if err != nil {
			c := color.New(color.FgRed, color.Bold).SprintFunc()
			p := c("[ERROR] ")
			fmt.Println(p, "Command failed to execute")
		}
	*/

	c := color.New(color.FgBlue, color.Bold).SprintFunc()
	p := c("[INFO] ")
	fmt.Println(p, "Command successfully executed.")

}
