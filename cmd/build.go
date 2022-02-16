package cmd

import (
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
	buildConfig, flashConfig := configparser.Parse()
	utils.DownloadFirmware()

	log.Print("Replacing firmware files with provided variables...")
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/nerdlight-firmware.ino", "{{VARS_AWS_ENDPOINT}}", buildConfig.AWSIOTEndpoint)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/nerdlight-firmware.ino", "{{VARS_R_PIN}}", buildConfig.RPin)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/nerdlight-firmware.ino", "{{VARS_G_PIN}}", buildConfig.GPin)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/nerdlight-firmware.ino", "{{VARS_B_PIN}}", buildConfig.BPin)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/Makefile", "{{core}}", flashConfig.ChipCore)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/Makefile", "{{chip}}", flashConfig.ChipType)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/Makefile", "{{boardconfig}}", flashConfig.ArduinoCliBuildBoardConfig)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/Makefile", "{{mkspiffs_bin}}", flashConfig.MKSPIFFSBinPath)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/Makefile", "{{device_port}}", flashConfig.DevicePort)
	utils.ReplaceInFile(buildConfig.FirmwareTempDir+"/Makefile", "{{spiffs_start_hex}}", flashConfig.SPIFFSStartHEX)
	log.Print("Firmware files ready for compilation!")

	log.Print("Loading certificates...")
	utils.CopyFile(buildConfig.CACertFilePath, buildConfig.FirmwareTempDir+"/data/ca.der")
	utils.CopyFile(buildConfig.PrivKeyPath, buildConfig.FirmwareTempDir+"/data/private.der")
	utils.CopyFile(buildConfig.CertPath, buildConfig.FirmwareTempDir+"/data/cert.der")

	log.Print("Starting firmware compilation")
	log.Print("Installing board...")
	errInitializingBoard := utils.LaunchCommand(buildConfig.FirmwareTempDir, "arduino-cli", "--additional-urls", "http://arduino.esp8266.com/stable/package_esp8266com_index.json", "core", "install", "esp8266:esp8266@2.7.4")
	if errInitializingBoard != nil {
		log.Fatal("[Board initialiation] Failed to compile firmware: ", errInitializingBoard)
		return
	}

	fwDeps := []string{
		"https://github.com/tzapu/WiFiManager.git",
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

	log.Print("Assembling FS...")
	errAssemblingFS := utils.LaunchCommand(buildConfig.FirmwareTempDir, "make", "filesystem.bin")
	if errAssemblingFS != nil {
		log.Fatal("[FS assembly] Failed to compile firmware: ", errAssemblingFS)
		return
	}

	log.Print("Compiling firmware...")
	err := utils.LaunchCommand(
		buildConfig.FirmwareTempDir,
		"make", "build",
	)
	if err != nil {
		log.Fatal("Failed to compile firmware: ", err)
		return
	}

	log.Printf("Compilation ended!")
}
