package cmd

import (
	"fmt"
	"github.com/forsam-education/kerberos/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// resetCmd represents the create command
var resetCmd = &cobra.Command{
	Use:   "reset-config",
	Short: "Resets the configuration to default values. WARNING: WILL OVERWRITE YOUR CONFIGURATION FILE",
	Long: `This command replaces your actual configuration file with the application default configuration values.

WARNING: WILL OVERWRITE YOUR CONFIGURATION FILE.`,
	Run: func(cmd *cobra.Command, args []string) {
		utils.SetConfigDefaults(true)
		utils.WriteConfig()
		fmt.Printf("%s\n\nConfiguration has been reset to default values in file %s\n", utils.ASCIILogo, viper.ConfigFileUsed())
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
