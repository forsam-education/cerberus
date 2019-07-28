package state

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d", viper.GetString(utils.RedisServerHost), viper.GetInt(utils.RedisServerPost)),
		DB: viper.GetInt(utils.RedisServerDBID),
	})
}