package configparser

import t "github.com/ziomarco/nerdlight-firmwarebuilder/cmd/types"

func Parse() (t.BuildConfig, t.FlashConfig) {
	// TODO: Implement
	// Viper also parses it, get values from that!
	return (t.BuildConfig{
			FirmwareGitUrl:    "https://github.com/ziomarco/nerdlight-firmware.git",
			FirmwareTempDir:   "/tmp/nerdlight-firmware",
			FirmwareGitBranch: "cli",
			AWSIOTEndpoint:    "a3l1l7s5ulqonx-ats.iot.eu-west-1.amazonaws.com",
			RPin:              "D1",
			GPin:              "D2",
			BPin:              "D6",
			CACertFilePath:    "/Users/ziomarco/Documents/Projects/nerdlight/firmware_built_certs/nodemcu_original/ca.der",
			PrivKeyPath:       "/Users/ziomarco/Documents/Projects/nerdlight/firmware_built_certs/nodemcu_original/private.der",
			CertPath:          "/Users/ziomarco/Documents/Projects/nerdlight/firmware_built_certs/nodemcu_original/cert.der",
		}), (t.FlashConfig{
			ChipCore:                   "esp8266:esp8266",
			ChipType:                   "esp8266",
			ArduinoCliBuildBoardConfig: "esp8266:esp8266:nodemcu",
			MKSPIFFSBinPath:            "~/Library/Arduino15/packages/esp8266/tools/mkspiffs/2.5.0-4-b40a506/mkspiffs",
			DevicePort:                 "/dev/cu.usbserial-0001",
			SPIFFSStartHEX:             "0x200000",
		})
}

func Write(buildcfg t.BuildConfig, flashcfg t.FlashConfig, configpath string) (status bool, err error) {
	// TODO: Implement
	return true, nil
}
