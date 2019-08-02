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

	nodeID, err := uuid.NewV4()
	if err != nil {
		return err
	}
	nodeIDString := nodeID.String()
	utils.Logger.Info(fmt.Sprintf("Node ID: %s", nodeIDString), nil)

	Manager = &manager{
		RedisClient:           redisClient,
		LeaderID:              nodeIDString,
		LeaderLockRefreshTime: viper.GetDuration(utils.LeaderLockRefreshTime) * time.Second,
	}

	if !Manager.IsRedisInitialized() {
		err := Manager.SetDefaultRedisState()
		if err == nil {
			go leaderRoutine()
		}

		return err
	}
	err = Manager.AddNode()
	if err != nil {
		return err
	}
	utils.Logger.Info("Successfully registered node into Redis.", nil)

	tryToGetLead()
	go leaderRoutine()

	return nil
}
