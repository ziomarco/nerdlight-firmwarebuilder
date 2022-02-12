package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type ConfigMap struct {
	OutputPath string
	Hash       string
}

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "initcfg",
	Short: "Create your configuration file",
	Long: `This command can be used to customize your configuration file.

	Simply follow the prompts and you'll be done!`,
	Run: _handleinit,
}

func _handleinit(cmd *cobra.Command, args []string) {
	// TODO: Implement
	fmt.Println("init called")
}

func init() {
	rootCmd.AddCommand(initCmd)
}
