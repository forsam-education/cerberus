package cmd

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use:   "cerberus",
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
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is $HOME/.cerberus/config.yaml)")
	rootCmd.PersistentFlags().BoolVarP(&utils.VerboseFlag, "verbose", "v", false, "Set cerberus to verbose mode")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := homedir.Dir()
		if err != nil {
			utils.Logger.StdErrorCritical(err, nil)
		}

		utils.SetConfigDefaults(false)

		viper.SetConfigType("toml")
		viper.AddConfigPath(home)
		viper.AddConfigPath(path.Join(home, ".cerberus"))
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
	}
	viper.SetEnvPrefix("cerberus")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	if !utils.FileExists(viper.ConfigFileUsed()) {
		utils.SetConfigDefaults(false)
		utils.WriteConfig()
	}

	viper.AutomaticEnv()

	_ = viper.ReadInConfig()
}
