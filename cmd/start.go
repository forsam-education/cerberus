package cmd

import (
	"fmt"
	"github.com/forsam-education/kerberos/api"
	"github.com/forsam-education/kerberos/proxy"
	"github.com/forsam-education/kerberos/utils"
	"github.com/spf13/cobra"
	"sync"
)

// startCmd represents the config command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the Kerberos reverse proxy.",
	Long: fmt.Sprintf(`%s

  A simple but powerful reverse proxy.

  Starts the Kerberos reverse proxy.`, utils.ASCIILogo),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(utils.ASCIILogo)

		var waitgroup sync.WaitGroup

		waitgroup.Add(2)
		go proxy.StartServer(&waitgroup)
		go api.StartServer(&waitgroup)
		waitgroup.Wait()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
