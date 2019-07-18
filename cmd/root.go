package cmd

import (
	"fmt"
	"github.com/forsam-education/kerberos/utils"
	"github.com/spf13/cobra"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "kerberos",
	Short: "A simple but powerful reverse proxy.",
	Long: fmt.Sprintf(`%s

  A simple but powerful reverse proxy.`, utils.ASCIILogo),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kerberos/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&utils.VerboseFlag, "verbose", "v", false, "Set kerberos to verbose mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			utils.LogAndForceExit(err)
		}

		utils.SetConfigDefaults(false)

		viper.SetConfigType("toml")
		viper.SetConfigFile(path.Join(home, ".kerberos/config.toml"))
	}
	viper.AutomaticEnv()

	if !utils.FileExists(viper.ConfigFileUsed()) {
		utils.SetConfigDefaults(true)
		utils.WriteConfig()
	}

	_ = viper.ReadInConfig()
}
