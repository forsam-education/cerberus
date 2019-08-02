package database

import (
	"fmt"
	"github.com/forsam-education/cerberus/utils"
	"github.com/spf13/viper"
)

func buildDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		viper.GetString(utils.DatabaseServerUser),
		viper.GetString(utils.DatabaseServerPass),
		viper.GetString(utils.DatabaseServerHost),
		viper.GetInt(utils.DatabaseServerPort),
		viper.GetString(utils.DatabaseServerDBName),
	)
}
