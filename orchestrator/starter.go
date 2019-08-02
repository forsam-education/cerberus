package orchestrator

import (
	"github.com/forsam-education/cerberus/database"
	"github.com/forsam-education/cerberus/proxy"
	"github.com/forsam-education/cerberus/state"
	"github.com/forsam-education/cerberus/utils"
	"github.com/volatiletech/sqlboiler/boil"
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

	if utils.IsLeaderNode {
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
		utils.LogAndForceExit(err)
	}

	err = handleDatabaseStartup()
	if err != nil {
		utils.LogAndForceExit(err)
	}

	if err = proxy.LoadServices(); err != nil {
		utils.LogAndForceExit(err)
	}

	var waitgroup sync.WaitGroup
	waitgroup.Add(1)

	go proxy.StartServer(&waitgroup)
	//go administration.StartServer(ctx, &waitgroup)

	waitgroup.Wait()

	shutdownOrchestrator()
}

func shutdownOrchestrator() {
	if utils.SharedStateManager != nil {
		if err := utils.SharedStateManager.Shutdown(); err != nil {
			utils.Logger.Critical(err.Error(), nil)
		}
	}
}
