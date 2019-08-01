package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"os"
)

// ConfigSetter should set a config to a given value.
type ConfigSetter func(key string, value interface{})

const (
	// ProxyServerHost is the host name for the Cerberus reverse proxy.
	ProxyServerHost = "http.proxy.host"
	// ProxyServerPort is the port for the Cerberus reverse proxy.
	ProxyServerPort = "http.proxy.port"
	// ProxyServerKeepAliveState allows to enable or diable HTTP Keep Alive on the Cerberus reverse proxy.
	ProxyServerKeepAliveState = "http.proxy.enable_keep_alive"
	// ProxyServerKeepAlivePeriod is the period in seconds to leave Keep Alive connections open.
	ProxyServerKeepAlivePeriod = "http.proxy.keep_alive_period"
	// ProxyServerReadTimeout is the period in seconds without read before timing out connection.
	ProxyServerReadTimeout = "http.proxy.read_timeout"
	// AdministrationServerHost is the host name for the Cerberus reverse proxy administration.
	AdministrationServerHost = "http.administration.host"
	// AdministrationServerPort is the port for the Cerberus reverse proxy administration.
	AdministrationServerPort = "http.administration.port"
	// DatabaseServerHost is the host name for the Cerberus MySQL database server.
	DatabaseServerHost = "database.host"
	// DatabaseServerPort is the port for the the Cerberus MySQL database server.
	DatabaseServerPort = "database.port"
	// DatabaseServerUser is the username for the the Cerberus MySQL database server.
	DatabaseServerUser = "database.user"
	// DatabaseServerPass is the password for the the Cerberus MySQL database server.
	DatabaseServerPass = "database.pass"
	// DatabaseServerDBName is the name of the database for the the Cerberus MySQL database server.
	DatabaseServerDBName = "database.dbname"
	// PasswordHashCost is the cost of the password hashing algorithm (BCrypt)
	PasswordHashCost = "misc.password_hash_cost"
)

// ASCIILogo is the ascii representation of the Athena logo
var ASCIILogo = `
   _____  ______  _____   ____   ______  _____   _    _   _____ 
  / ____||  ____||  __ \ |  _ \ |  ____||  __ \ | |  | | / ____|
 | |     | |__   | |__) || |_) || |__   | |__) || |  | || (___  
 | |     |  __|  |  _  / |  _ < |  __|  |  _  / | |  | | \___ \ 
 | |____ | |____ | | \ \ | |_) || |____ | | \ \ | |__| | ____) |
  \_____||______||_|  \_\|____/ |______||_|  \_\ \____/ |_____/

`

// SetConfigDefaults resets the configuration to the default value
func SetConfigDefaults(force bool) {
	var setConfig ConfigSetter
	if force {
		setConfig = viper.Set
	} else {
		setConfig = viper.SetDefault
	}
	setConfig(ProxyServerHost, "127.0.0.1")
	setConfig(ProxyServerPort, 8970)
	setConfig(ProxyServerKeepAliveState, true)
	setConfig(ProxyServerKeepAlivePeriod, 5)
	setConfig(ProxyServerReadTimeout, 5)
	setConfig(AdministrationServerHost, "127.0.0.1")
	setConfig(AdministrationServerPort, 8971)
	setConfig(DatabaseServerHost, "database")
	setConfig(DatabaseServerPort, 3306)
	setConfig(DatabaseServerUser, "root")
	setConfig(DatabaseServerPass, "root")
	setConfig(DatabaseServerDBName, "cerberus")
	setConfig(PasswordHashCost, 15)
}

// WriteConfig replaces the config file by the current configuration.
func WriteConfig() {
	filePath := viper.ConfigFileUsed()
	if filePath == "" {
		filePath = "./config.toml"
	}
	_ = os.Remove(filePath)
	CreateFile(filePath)
	err := viper.WriteConfig()
	if nil != err {
		log.Fatalln(err)
	}
}

// BuildDBDSN returns a complete MySQL DSN from configuration.
func BuildDBDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?parseTime=true",
		viper.GetString(DatabaseServerUser),
		viper.GetString(DatabaseServerPass),
		viper.GetString(DatabaseServerHost),
		viper.GetInt(DatabaseServerPort),
		viper.GetString(DatabaseServerDBName),
	)
}
