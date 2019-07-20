package proxy

import (
	"fmt"
	"github.com/forsam-education/kerberos/models"
	"github.com/forsam-education/kerberos/utils"
)

var services models.ServiceSlice

func LoadServices() error {
	utils.Logger.Info("Loading services from database...", nil)

	var err error
	services, err = models.Services().AllG()
	if err != nil {
		return err
	}

	for _, service := range services {
		utils.Logger.Info(fmt.Sprintf("%s - %s - %s", service.Name, service.Path, service.TargetURL), nil)
	}

	return nil
}
