package state

import (
	"encoding/json"
	"fmt"
	"github.com/forsam-education/cerberus/models"
	"github.com/forsam-education/cerberus/utils"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"time"
)

const (
	// CurrentNodesCount is the redis key for current Cerberus node count.
	CurrentNodesCount = "nodecount"
	// CurrentRequestsCount is the redis key for current Cerberus proxy request count.
	CurrentRequestsCount = "requestcount"
	// LeaderLock is the lock to get leadership.
	LeaderLock = "leaderlock"
)

// Manager is the state manager for Cerberus.
var Manager *manager

type manager struct {
	RedisClient           *redis.Client
	Registered            bool
	LeaderID              string
	LeaderLockRefreshTime time.Duration
	IsLeaderNode          bool
}

// GetCurrentRequestsCount fetches the current proxy request from the shared manager.
func (manager *manager) GetCurrentRequestsCount() (int, error) {
	return manager.RedisClient.Get(CurrentRequestsCount).Int()
}

// AddRequest adds a request to the current request count in the shared manager.
func (manager *manager) AddRequest() error {
	return manager.RedisClient.Incr(CurrentRequestsCount).Err()
}

// RemoveRequest removes a request from the current request count in the shared manager.
func (manager *manager) RemoveRequest() error {
	return manager.RedisClient.Decr(CurrentRequestsCount).Err()
}

// GetAndResetRequestCount performs a getset on the manager to get the request count and reset it to 0 in an atomic operation.
func (manager *manager) GetAndResetRequestCount() (int, error) {
	return manager.RedisClient.GetSet(CurrentRequestsCount, 0).Int()
}

// GetCurrentNodesCount fetches the cerberus nodes from the shared manager.
func (manager *manager) GetCurrentNodesCount() (int, error) {
	return manager.RedisClient.Get(CurrentNodesCount).Int()
}

// AddNode registers the current node in the shared manager.
func (manager *manager) AddNode() error {
	if err := manager.RedisClient.Incr(CurrentNodesCount).Err(); err != nil {
		return err
	}
	manager.Registered = true

	return nil
}

// RemoveNode deregisters the current node the shared manager.
func (manager *manager) RemoveNode() error {
	if err := manager.RedisClient.Decr(CurrentNodesCount).Err(); err != nil {
		return err
	}
	manager.Registered = false

	return nil
}

// Shutdown checks if the node is registers and deregisters it from the shared manager.
func (manager *manager) Shutdown() error {
	if !manager.Registered {
		return nil
	}
	if err := manager.RemoveNode(); err != nil {
		return fmt.Errorf("can't deregister node from Redis: %s", err.Error())
	}

	utils.Logger.Info("Successfully deregistered node from Redis.", nil)

	return nil
}

// IsRedisInitialized checks if there is any node in the cerberus cluster.
func (manager *manager) IsRedisInitialized() bool {
	count, _ := manager.GetCurrentNodesCount()

	return count > 0
}

// SetDefaultRedisState sets the default state in the shared manager and registers the current node.
func (manager *manager) SetDefaultRedisState() error {
	// Add current node
	if err := manager.AddNode(); err != nil {
		return err
	}
	utils.Logger.Info("Successfully registered node into Redis.", nil)

	// Set current request count to 0
	if err := manager.RedisClient.Set(CurrentRequestsCount, 0, 0).Err(); err != nil {
		return err
	}
	utils.Logger.Info("Successfully set default Redis state.", nil)

	return nil
}

// TryToAcquireLead tries to acquire a lock in the Redis DB and set the leader status accordingly.
func (manager *manager) TryToAcquireLead() bool {
	wasUnset, err := manager.RedisClient.SetNX(LeaderLock, manager.LeaderID, manager.LeaderLockRefreshTime).Result()
	if err != nil {
		manager.IsLeaderNode = false
		return manager.IsLeaderNode
	}
	lockID, err := manager.RedisClient.Get(LeaderLock).Result()
	if err != nil || lockID != manager.LeaderID {
		manager.IsLeaderNode = false
		return manager.IsLeaderNode
	}
	if !wasUnset {
		manager.RedisClient.Expire(LeaderLock, manager.LeaderLockRefreshTime)
	}

	manager.IsLeaderNode = true
	return manager.IsLeaderNode
}

func (manager *manager) FindServiceByPath(path []byte) (*models.Service, error) {
	var service *models.Service
	result, err := manager.RedisClient.Get(string(path)).Bytes()

	if err != nil {
		// Weird behavior of Redis lib
		// Nil result is considered as an untyped error, we must check for error message
		// Also note that their Nil type is internal and this may change in the future
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}

	err = json.Unmarshal(result, service)
	if err != nil {
		return nil, err
	}

	return service, nil
}

func (manager *manager) AddService(service *models.Service) error {
	json, err := json.Marshal(service)
	if err != nil {
		return err
	}

	return manager.RedisClient.Set(service.ServicePath, string(json), viper.GetDuration(utils.RedisServerServiceTTL)*time.Minute).Err()
}
