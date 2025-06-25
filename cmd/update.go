package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates mysqldumpmanager based on the configuration file",
	Long: `Updates mysqldumpmanager based on the configuration file. 

By default, the configuration file is located at /etc/mysqldumpmanager/conf.yaml`,
	Run: updateCmdFunc,
}

func init() {
	rootCmd.AddCommand(updateCmd)
}

func updateCmdFunc(cmd *cobra.Command, args []string) {

	c := color.New(color.FgBlue, color.Bold).SprintFunc()
	p := c("[INFO] ")
	fmt.Println(p, "Looking for conf.yaml and updating MySQLDumpManager")

}

func fetchYamlConfig() {

}

func pushYamlConfig() {

}
