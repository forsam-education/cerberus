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

const (
	CurrentNodesCount    = "nodecount"
	CurrentRequestsCount = "requestcount"
)

func isRedisInitialized() bool {
	count, _ := Manager.GetCurrentNodesCount()

	return count > 0
}

func setDefaultState() error {
	// Add current node
	if err := Manager.AddNode(); err != nil {
		return err
	}

	// Set current request count to 0
	if err := Manager.redisClient.Set(CurrentRequestsCount, 0, 0).Err(); err != nil {
		return err
	}

	return nil
}

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

	Manager = &manager{
		redisClient: redisClient,
	}

	if !isRedisInitialized() {
		err := setDefaultState()
		return err
	}
	err := Manager.AddNode()

	return err
}
