package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Deletes selected dump",
	Long: `Deletes selected dump. 

By default, dumps are located in /etc/mysqldumpmanager/dumps/`,
	Run: deleteCmdFunc,
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("file", "F", "sample.sql", "file to be deleted")
}

func deleteCmdFunc(cmd *cobra.Command, args []string) {
	file, err := cmd.Flags().GetString("file")
	if err != nil {
		c := color.New(color.FgRed, color.Bold).SprintFunc()
		p := c("[ERROR] ")
		fmt.Println(p, "Command failed to execute")
	}

	c := color.New(color.FgBlue, color.Bold).SprintFunc()
	p := c("[INFO] ")
	fmt.Println(p, file, " Command successfully executed.")

}
