package orchestrator

import (
	"context"
	"github.com/forsam-education/cerberus/administration"
	"github.com/forsam-education/cerberus/database"
	"github.com/forsam-education/cerberus/proxy"
	"github.com/forsam-education/cerberus/state"
	"github.com/forsam-education/cerberus/utils"
	"github.com/volatiletech/sqlboiler/boil"
	"os"
	"sync"
)

func handleDatabaseStartup() error {
	db, err := database.Connect()
	if err != nil {
		return err
	}
	utils.Logger.Info("Connected to database.", nil)

	err = database.Migrate(db)
	if err != nil {
		return err
	}

	boil.SetDB(db)

	if state.Manager.IsLeaderNode {
		if err = database.HandleFirstStart(); err != nil {
			return err
		}
	}

	return nil
}

// StartOrchestrator is the hearth of the machine, it orchestrates services and managers.
func StartOrchestrator() {
	err := state.InitManager()
	if err != nil {
		utils.Logger.StdErrorCritical(err, nil)
		os.Exit(1)
	}

	err = handleDatabaseStartup()
	if err != nil {
		utils.Logger.StdErrorCritical(err, nil)
		_ = state.Manager.RemoveNode()
		os.Exit(1)
	}

	if err = proxy.LoadServices(); err != nil {
		utils.Logger.StdErrorCritical(err, nil)
		_ = state.Manager.RemoveNode()
		os.Exit(1)
	}

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	go proxy.StartServer(&waitgroup)
	go administration.StartServer(context.Background(), &waitgroup)

	waitgroup.Wait()

	shutdownOrchestrator()
}

func shutdownOrchestrator() {
	if state.Manager != nil {
		if err := state.Manager.Shutdown(); err != nil {
			utils.Logger.Critical(err.Error(), nil)
		}
	}
}
