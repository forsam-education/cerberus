package cmd

import (
	"fmt"
	"github.com/forsam-education/cerberus/orchestrator"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the Cerberus reverse proxy.",
	Long: fmt.Sprintf(`%s

  A simple but powerful reverse proxy.

  Starts the Cerberus reverse proxy.`, utils.ASCIILogo),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(utils.ASCIILogo)

		orchestrator.StartOrchestrator()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
