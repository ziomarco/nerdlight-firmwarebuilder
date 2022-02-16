package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/ziomarco/nerdlight-firmwarebuilder/cmd/configparser"
	"github.com/ziomarco/nerdlight-firmwarebuilder/cmd/utils"
)

var flashCmd = &cobra.Command{
	Use:   "flash",
	Short: "Use this command to flash your device",
	Long: `With this command you'll be able to flash
nerdlight firmware on your device.

Usage: nerdlight-fb flash`,
	Run: _handleFlash,
}

func init() {
	rootCmd.AddCommand(flashCmd)
}

func _handleFlash(cmd *cobra.Command, args []string) {
	buildConfig, _ := configparser.Parse()

	log.Print("Flashing FS...")
	errFlashingFS := utils.LaunchCommand(
		buildConfig.FirmwareTempDir,
		"make", "flash-fs",
	)
	if errFlashingFS != nil {
		log.Fatal("Failed to flash SPIFFS FS: ", errFlashingFS)
		return
	}

	log.Print("Flashing firmware...")
	err := utils.LaunchCommand(
		buildConfig.FirmwareTempDir,
		"make", "flash",
	)
	if err != nil {
		log.Fatal("Failed to flash firmware: ", err)
		return
	}

	log.Printf("Flash ended!")
}
