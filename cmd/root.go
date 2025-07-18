/*
This is the main entry point for the cobra library and serves as the CLI client manager
*/

package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "mysqldumpmanager",
	Version: "0.1 (Pre-Alpha release)",
	Short:   "A lightweight application to create, list and delete MySQL backups",
	Long: `MySQLDumpManager is a CLI tool to create, list and delete MySQL backups through the widely used out-of-the-box client utility mysqldump.

Further information regarding mysqldump's functionality go can be read on the official documentation at: https://dev.mysql.com/doc/refman/8.4/en/mysqldump.html`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Disabling the default Completion command as per cobra documentation
	rootCmd.CompletionOptions.DisableDefaultCmd = true

}
