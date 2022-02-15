package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/ziomarco/nerdlight-firmwarebuilder/cmd/configparser"
	"github.com/ziomarco/nerdlight-firmwarebuilder/cmd/utils"
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
	buildConfig, _ := configparser.Parse()
	utils.DownloadFirmware()

	log.Print("Replacing firmware files with provided variables...")
	utils.ReplaceInFile("/tmp/nerdlight-firmware/nerdlight-firmware.ino", "{{VARS_AWS_ENDPOINT}}", buildConfig.AWSIOTEndpoint)
	utils.ReplaceInFile("/tmp/nerdlight-firmware/nerdlight-firmware.ino", "{{VARS_R_PIN}}", buildConfig.RPin)
	utils.ReplaceInFile("/tmp/nerdlight-firmware/nerdlight-firmware.ino", "{{VARS_G_PIN}}", buildConfig.GPin)
	utils.ReplaceInFile("/tmp/nerdlight-firmware/nerdlight-firmware.ino", "{{VARS_B_PIN}}", buildConfig.BPin)
	log.Print("Firmware files ready for compilation!")

	log.Print("Starting firmware compilation")
	log.Print("Installing board...")
	errInitializingBoard := utils.LaunchCommand(buildConfig.FirmwareTempDir, "arduino-cli", "--additional-urls", "http://arduino.esp8266.com/stable/package_esp8266com_index.json", "core", "install", "esp8266:esp8266@2.7.4")
	if errInitializingBoard != nil {
		log.Fatal("[Board initialiation] Failed to compile firmware: ", errInitializingBoard)
		return
	}

	fwDeps := []string{
		"https://github.com/tzapu/WiFiManager.git",
		"https://github.com/esp8266/Arduino.git",
		"https://github.com/knolleary/pubsubclient.git",
		"https://github.com/arduino-libraries/NTPClient.git",
	}
	for idx, dep := range fwDeps {
		log.Printf("Installing deps %v of %v", idx+1, len(fwDeps))
		errInitializingPlugins := utils.LaunchCommandWithEnv(
			buildConfig.FirmwareTempDir,
			[]string{"ARDUINO_LIBRARY_ENABLE_UNSAFE_INSTALL=true"},
			"arduino-cli", "lib", "install", "--git-url", dep,
		)
		if errInitializingPlugins != nil {
			log.Fatal("[Plugins initialization] Failed to compile firmware: ", errInitializingPlugins)
			return
		}
	}

	const boardFQDN string = "esp8266:esp8266:espduino"
	const firmwareSketchFile string = "nerdlight-firmware.ino"
	log.Print("Compiling firmware...")
	err := utils.LaunchCommand(buildConfig.FirmwareTempDir, fmt.Sprintf("arduino-cli compile --fqbn %s %s", boardFQDN, firmwareSketchFile))
	if err != nil {
		log.Fatal("Failed to compile firmware: ", err)
		return
	}

	log.Printf("Compilation ended! Your built firmware is in: %s", buildConfig.FirmwareTempDir+"/output")
}
