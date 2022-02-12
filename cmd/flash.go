package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var flashCmd = &cobra.Command{
	Use:   "flash",
	Short: "Use this command to flash your device",
	Long: `With this command you'll be able to flash
nerdlight firmware on your device.

Usage: nerdlight-fb flash <options>`,
	Run: _handleFlash,
}

func init() {
	rootCmd.AddCommand(flashCmd)
	flashCmd.PersistentFlags().String("device-port", "", "The serial port of your device (e.g. /dev/ttyUSB0)")
	flashCmd.PersistentFlags().String("firmware", "", "Complete path to your firmware binary file")
}

func _handleFlash(cmd *cobra.Command, args []string) {
	// TODO: Implement
	fmt.Println("flash called")
}
