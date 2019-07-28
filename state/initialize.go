package state

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/forsam-education/simplelogger"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// InitManager connects to redis and setup the state manager.
func InitManager() error {
	host := fmt.Sprintf("%s:%d", viper.GetString(utils.RedisServerHost), viper.GetInt(utils.RedisServerPort))

	// Catch interrupt signal in channel.
	signalCatcher := make(chan os.Signal, 1)
	signal.Notify(signalCatcher, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	utils.Logger.Info("Connecting to redis client...", simplelogger.LogExtraData{"host": host, "databaseId": viper.GetInt(utils.RedisServerDBID)})
	redisClient := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: viper.GetString(utils.RedisServerPass),
		DB:       viper.GetInt(utils.RedisServerDBID),
	})

	if err := redisClient.Ping().Err(); err != nil {
		return err
	}

	utils.SharedStateManager = &utils.StateManager{
		RedisClient: redisClient,
	}

	if !utils.SharedStateManager.IsRedisInitialized() {
		err := utils.SharedStateManager.SetDefaultRedisState()
		return err
	}
	err := utils.SharedStateManager.AddNode()
	utils.Logger.Info("Successfully registered node into Redis.", nil)

	return err
}
