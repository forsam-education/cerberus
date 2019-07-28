package state

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/go-redis/redis"
)

type manager struct {
	redisClient *redis.Client
}

// Manager is the shared state manager for Cerberus.
var Manager *manager

func (manager *manager) GetCurrentRequestsCount() (int, error) {
	return manager.redisClient.Get(CurrentRequestsCount).Int()
}

func (manager *manager) AddRequest() error {
	return manager.redisClient.Incr(CurrentRequestsCount).Err()
}

func (manager *manager) RemoveRequest() error {
	return manager.redisClient.Decr(CurrentRequestsCount).Err()
}

func (manager *manager) GetCurrentNodesCount() (int, error) {
	return manager.redisClient.Get(CurrentNodesCount).Int()
}

func (manager *manager) AddNode() error {
	return manager.redisClient.Incr(CurrentNodesCount).Err()
}

func (manager *manager) RemoveNode() error {
	return manager.redisClient.Decr(CurrentNodesCount).Err()
}

func (manager *manager) Shutdown() error {
	if err := manager.RemoveNode(); err != nil {
		return fmt.Errorf("can't deregister node from Redis: %s", err.Error())
	}

	utils.Logger.Info("Successfully deregistered node from Redis.", nil)

	return nil
}
