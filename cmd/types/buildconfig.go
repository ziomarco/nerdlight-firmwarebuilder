package types

type BuildConfig struct {
	// TODO: Define
	FirmwareTempDir   string
	FirmwareGitUrl    string
	FirmwareGitBranch string
	AWSIOTEndpoint    string
	RPin              string
	GPin              string
	BPin              string
	CACertFilePath    string
	PrivKeyPath       string
	CertPath          string
}
