package cmd

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/forsam-education/cerberus/administration"
	"github.com/forsam-education/cerberus/database"
	"github.com/forsam-education/cerberus/proxy"
	"github.com/forsam-education/cerberus/state"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/cobra"
	"github.com/volatiletech/sqlboiler/boil"
	"os"
	"sync"
	"time"
)

// startCmd represents the config command.
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the Cerberus reverse proxy.",
	Long: fmt.Sprintf(`%s

  A simple but powerful reverse proxy.

  Starts the Cerberus reverse proxy.`, utils.ASCIILogo),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(utils.ASCIILogo)

		dsn := utils.BuildDBDSN()

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			utils.LogAndForceExit(err)
		}

		utils.Logger.Info("Connecting to database...", map[string]interface{}{"DSN": dsn})
		var dbErr error
		for i := 1; i <= 3; i++ {
			dbErr = db.Ping()
			if dbErr != nil {
				utils.Logger.Info(fmt.Sprintf("Attempt #%d failed, will retry in 10 seconds", i), map[string]interface{}{"Error": dbErr})
				time.Sleep(10 * time.Second)
				continue
			}

			break
		}

		if dbErr != nil {
			utils.Logger.Error("Can't connect to database after 3 attempts.", nil)
			os.Exit(1)
		}

		utils.Logger.Info("Connected to database.", nil)

		err = database.MigrateDatabase(db)
		if err != nil {
			utils.LogAndForceExit(err)
		}

		err = state.InitManager()
		if err != nil {
			utils.LogAndForceExit(err)
		}

		boil.SetDB(db)

		if err = database.HandleFirstStart(); err != nil {
			utils.LogAndForceExit(err)
		}

		if err = proxy.LoadServices(); err != nil {
			utils.LogAndForceExit(err)
		}

		ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

		var waitgroup sync.WaitGroup
		waitgroup.Add(2)
		go proxy.StartServer(ctx, &waitgroup)
		go administration.StartServer(ctx, &waitgroup)

		waitgroup.Wait()

		if err = state.Manager.Shutdown(); err != nil {
			utils.Logger.Critical(err.Error(), nil)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
