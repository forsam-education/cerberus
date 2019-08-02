package state

import (
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/viper"
	"time"
)

func tryToGetLead() {
	wasLeader := Manager.IsLeaderNode
	isLeader := Manager.TryToAcquireLead()
	if isLeader {
		if !wasLeader {
			utils.Logger.Info("Node is the leader.", nil)
		}
	} else {
		if wasLeader {
			utils.Logger.Info("Node is now a worker.", nil)
		}
	}
}

func leaderRoutine() {
	for {
		tryToGetLead()
		time.Sleep(viper.GetDuration(utils.LeaderLockRefreshTime)*time.Second - 2*time.Second)
	}
}
