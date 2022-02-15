package utils

import (
	"log"
	"os"

	"github.com/ziomarco/nerdlight-firmwarebuilder/cmd/configparser"
)

func DownloadFirmware() {
	log.Print("Downloading firmware...")
	buildConfig, _ := configparser.Parse()
	var giturl string = buildConfig.FirmwareGitUrl
	var outdir string = buildConfig.FirmwareTempDir
	var branch string = buildConfig.FirmwareGitBranch

	if _, err := os.Stat(outdir); !os.IsNotExist(err) {
		log.Printf("Removing old firmware directory: %s", outdir)
		os.RemoveAll(outdir)
	}

	LaunchCommand("", "git", "clone", giturl, outdir)
	LaunchCommand(outdir, "git", "checkout", branch)
	LaunchCommand(outdir, "git", "pull", "origin", branch)
	log.Println("Firmware downloaded!")
}
