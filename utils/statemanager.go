package utils

import (
	"fmt"
	"github.com/go-redis/redis"
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

// IsLeaderNode stores the leading state of the node.
var IsLeaderNode = false

// StateManager is the state manager type for Cerberus.
type StateManager struct {
	RedisClient           *redis.Client
	Registered            bool
	LeaderID              string
	LeaderLockRefreshTime time.Duration
}

// SharedStateManager is the shared state manager for Cerberus.
var SharedStateManager *StateManager

// GetCurrentRequestsCount fetches the current proxy request from the shared manager.
func (manager *StateManager) GetCurrentRequestsCount() (int, error) {
	return manager.RedisClient.Get(CurrentRequestsCount).Int()
}

// AddRequest adds a request to the current request count in the shared manager.
func (manager *StateManager) AddRequest() error {
	return manager.RedisClient.Incr(CurrentRequestsCount).Err()
}

// RemoveRequest removes a request from the current request count in the shared manager.
func (manager *StateManager) RemoveRequest() error {
	return manager.RedisClient.Decr(CurrentRequestsCount).Err()
}

// GetCurrentNodesCount fetches the cerberus nodes from the shared manager.
func (manager *StateManager) GetCurrentNodesCount() (int, error) {
	return manager.RedisClient.Get(CurrentNodesCount).Int()
}

// AddNode registers the current node in the shared manager.
func (manager *StateManager) AddNode() error {
	if err := manager.RedisClient.Incr(CurrentNodesCount).Err(); err != nil {
		return err
	}
	manager.Registered = true

	return nil
}

// RemoveNode deregisters the current node the shared manager.
func (manager *StateManager) RemoveNode() error {
	if err := manager.RedisClient.Decr(CurrentNodesCount).Err(); err != nil {
		return err
	}
	manager.Registered = false

	return nil
}

// Shutdown checks if the node is registers and deregisters it from the shared manager.
func (manager *StateManager) Shutdown() error {
	if !manager.Registered {
		return nil
	}
	if err := manager.RemoveNode(); err != nil {
		return fmt.Errorf("can't deregister node from Redis: %s", err.Error())
	}

	Logger.Info("Successfully deregistered node from Redis.", nil)

	return nil
}

// IsRedisInitialized checks if there is any node in the cerberus cluster.
func (manager *StateManager) IsRedisInitialized() bool {
	count, _ := manager.GetCurrentNodesCount()

	return count > 0
}

// SetDefaultRedisState sets the default state in the shared manager and registers the current node.
func (manager *StateManager) SetDefaultRedisState() error {
	// Add current node
	if err := SharedStateManager.AddNode(); err != nil {
		return err
	}
	Logger.Info("Successfully registered node into Redis.", nil)

	// Set current request count to 0
	if err := SharedStateManager.RedisClient.Set(CurrentRequestsCount, 0, 0).Err(); err != nil {
		return err
	}
	Logger.Info("Successfully set default Redis state.", nil)

	return nil
}

// TryToAcquireLead tries to acquire a lock in the Redis DB and set the leader status accordingly.
func (manager *StateManager) TryToAcquireLead() bool {
	wasUnset, err := manager.RedisClient.SetNX(LeaderLock, manager.LeaderID, manager.LeaderLockRefreshTime).Result()
	if err != nil {
		IsLeaderNode = false
		return IsLeaderNode
	}
	lockID, err := manager.RedisClient.Get(LeaderLock).Result()
	if err != nil || lockID != manager.LeaderID {
		IsLeaderNode = false
		return IsLeaderNode
	}
	if !wasUnset {
		manager.RedisClient.Expire(LeaderLock, manager.LeaderLockRefreshTime)
	}

	IsLeaderNode = true
	return IsLeaderNode
}
