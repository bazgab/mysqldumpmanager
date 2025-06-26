package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes MySQLDumpManager (first-time use)",
	Long: `Initializes MySQLDumpManager. This command is meant to be used only once after installation, 
and provides the instructions for MySQLDumpManager to create the necessary directories and files to make
it work as intended.`,
	Run: initCmdFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initCmdFunc(cmd *cobra.Command, args []string) {
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
