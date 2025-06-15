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
	Use:   "mysqldumpmanager",
	Short: "A lightweight application to create, list and delete MySQL backups",
	Long: `MySQLDumpManager is a CLI tool to create, list and delete MySQL backups through the widely used out-of-the-box client utility mysqldump.

Further information regarding mysqldump's functionality can be read on the official documentation at: https://dev.mysql.com/doc/refman/8.4/en/mysqldump.html`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.bum.yaml)")
}
