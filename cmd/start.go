package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/forsam-education/kerberos/api"
	"github.com/forsam-education/kerberos/database"
	"github.com/forsam-education/kerberos/proxy"
	"github.com/forsam-education/kerberos/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/boil"
	"sync"
	"time"
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

		dbDsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s",
			viper.GetString(utils.DatabaseServerUser),
			viper.GetString(utils.DatabaseServerPass),
			viper.GetString(utils.DatabaseServerHost),
			viper.GetInt(utils.DatabaseServerPort),
			viper.GetString(utils.DatabaseServerDBName),
		)

		db, err := sql.Open("mysql", dbDsn)
		if err != nil {
			utils.LogAndForceExit(err)
		}
		if err = db.Ping(); err != nil {
			utils.LogAndForceExit(err)
		}

		boil.SetDB(db)

		if err = database.HandleFirstStart(); err != nil {
			utils.LogAndForceExit(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		var waitgroup sync.WaitGroup

		waitgroup.Add(2)
		go proxy.StartServer(ctx, &waitgroup)
		go api.StartServer(ctx, &waitgroup)
		time.Sleep(4 * time.Second)
		waitgroup.Wait()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
