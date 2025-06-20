package cmd

import (
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
	loggingCmd.Flags().StringP("name", "N", "", "Input name to be displayed (required)")
	err := loggingCmd.MarkFlagRequired("name")
	if err != nil {
		return
	}

}

func loggingCmdFunc(cmd *cobra.Command, args []string) {
	name, err := cmd.Flags().GetString("name")
	if err != nil {
		m := "Failed to execute command " + name
		LogError(m)
	}
	m := "Command executed successfully"
	LogInfo(m)
	LogInfo("Logged: " + name)
}
