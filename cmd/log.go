package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var loggingCmd = &cobra.Command{
	Use:   "log",
	Short: "Log a word",
	Long:  "This logs a huge word",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: loggingCmdFunc,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.

func init() {
	rootCmd.AddCommand(loggingCmd)
	loggingCmd.Flags().StringP("name", "N", "hello", "Input name to be displayed")

}

func loggingCmdFunc(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		c := color.New(color.FgRed, color.Bold).SprintFunc()
		p := c("[ERROR] ")
		fmt.Println(p, "Command failed to execute")
	}
	m := "Command executed successfully"
	logInfo(m)
	logInfo("Logged: " + name)

}
