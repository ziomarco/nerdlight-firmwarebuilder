package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nerdlight-fb",
	Short: "Nerdlight NodeMCU Firmware Builder CLI ⚒",
	Long: `
                                 dP dP oo          dP         dP   
                                 88 88             88         88   
88d888b. .d8888b. 88d888b. .d888b88 88 dP .d8888b. 88d888b. d8888P 
88'  '88 88ooood8 88'  '88 88'  '88 88 88 88'  '88 88'  '88   88   
88    88 88.  ... 88       88.  .88 88 88 88.  .88 88    88   88   
dP    dP '88888P' dP       '88888P8 dP dP '8888P88 dP    dP   dP   
                                               .88                 
                                           d8888P
										   
This is a simple CLI written (bad) in Go for automating
ESP8266 NodeMCU firmware and flash to it!`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nerdlight-fb.yaml)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".nerdlight-fb" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".nerdlight-fb")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
