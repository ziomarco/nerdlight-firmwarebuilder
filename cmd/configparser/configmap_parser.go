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
	}), (t.FlashConfig{})
}

func Write(buildcfg t.BuildConfig, flashcfg t.FlashConfig, configpath string) (status bool, err error) {
	// TODO: Implement
	return true, nil
}
