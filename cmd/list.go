package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists current dumps",
	Long: `Lists current dumps. 

By default, they are located in /etc/mysqldumpmanager/dumps/`,
	Run: listCmdFunc,
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func listCmdFunc(cmd *cobra.Command, args []string) {
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
