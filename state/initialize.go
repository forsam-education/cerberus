package state

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/forsam-education/simplelogger"
	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func startLeadProcess() {
	for {
		wasLeader := utils.IsLeaderNode
		isLeader := utils.SharedStateManager.TryToAcquireLead()
		if isLeader {
			if !wasLeader {
				utils.Logger.Info("Node is the leader.", nil)
			}
		} else {
			if wasLeader {
				utils.Logger.Info("Node is now a worker.", nil)
			}
		}
		time.Sleep(8 * time.Second)
	}
}

// InitManager connects to redis and setup the state manager.
func InitManager() error {
	host := fmt.Sprintf("%s:%d", viper.GetString(utils.RedisServerHost), viper.GetInt(utils.RedisServerPort))

	// Catch interrupt signal in channel.
	signalCatcher := make(chan os.Signal, 1)
	signal.Notify(signalCatcher, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	utils.Logger.Info("Connecting to redis server...", simplelogger.LogExtraData{"host": host, "databaseId": viper.GetInt(utils.RedisServerDBID)})
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: viper.GetString(utils.RedisServerPass),
		DB:       viper.GetInt(utils.RedisServerDBID),
	})

	if err := redisClient.Ping().Err(); err != nil {
		return err
	}

	nodeId, err := uuid.NewV4()
	if err != nil {
		return err
	}
	nodeIdString := nodeId.String()
	utils.Logger.Info(fmt.Sprintf("Node ID: %s", nodeIdString), nil)

	utils.SharedStateManager = &utils.StateManager{
		RedisClient: redisClient,
		LeaderId:    nodeIdString,
	}

	if !utils.SharedStateManager.IsRedisInitialized() {
		err := utils.SharedStateManager.SetDefaultRedisState()
		if err == nil {
			go startLeadProcess()
		}

		return err
	}
	err = utils.SharedStateManager.AddNode()
	if err != nil {
		return err
	}
	utils.Logger.Info("Successfully registered node into Redis.", nil)

	go startLeadProcess()

	return nil
}
