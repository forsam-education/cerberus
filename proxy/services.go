package proxy

import (
	"fmt"
	"github.com/forsam-education/cerberus/models"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/viper"
	"time"
)

// MemoryCachedService represents a service cached in hot memory with an expiration.
type MemoryCachedService struct {
	Service        *models.Service
	ExpirationTime time.Time
}

// IsExpired returns true if the cache has expired, false otherwise.
func (service *MemoryCachedService) IsExpired() bool {
	return service.ExpirationTime.After(time.Now())
}

var services = make(map[string]MemoryCachedService)

// LoadServices gets services from the database and sets them to the services variable.
func LoadServices() error {
	utils.Logger.Info("Loading services from database...", nil)

	rawServices, err := models.Services().AllG()

	if err != nil {
		return err
	}

	for _, service := range rawServices {
		utils.Logger.Info(fmt.Sprintf("%s - %s - %s", service.Name, service.ServicePath, service.TargetHost), nil)
		services[service.ServicePath] = MemoryCachedService{
			Service:        service,
			ExpirationTime: time.Now().Add(viper.GetDuration(utils.ServiceHotMemoryTTL) * time.Second),
		}
	}

	return nil
}

func cacheService(service *models.Service) {
	services[service.ServicePath] = MemoryCachedService{
		Service:        service,
		ExpirationTime: time.Now().Add(viper.GetDuration(utils.ServiceHotMemoryTTL) * time.Second),
	}
}
