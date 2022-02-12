package cmd

import (
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Use this command to build your firmware",
	Long: `With this command you can build your firmware following
customized with configuration file parameters, ready to flash on your
device!`,
	Run: _handleBuild,
}

func init() {
	rootCmd.AddCommand(buildCmd)
	// buildCmd.PersistentFlags().String("foo", "", "A help for foo")
}

func _handleBuild(cmd *cobra.Command, args []string) {
	// TODO: Implement
}
