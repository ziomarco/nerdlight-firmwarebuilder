package configparser

import t "github.com/ziomarco/nerdlight-firmwarebuilder/cmd/types"

func Parse() (t.BuildConfig, t.FlashConfig) {
	// TODO: Implement
	// Viper also parses it, get values from that!
	return (t.BuildConfig{}), (t.FlashConfig{})
}

func Write(buildcfg t.BuildConfig, flashcfg t.FlashConfig, configpath string) (status bool, err error) {
	// TODO: Implement
	return true, nil
}
