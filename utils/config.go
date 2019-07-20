package utils

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

// VerboseFlag describes if Kerberos is in verbose mode.
var VerboseFlag bool

// ConfigSetter should set a config to a given value.
type ConfigSetter func(key string, value interface{})

const (
	// ProxyServerHost is the host name for the Kerberos reverse proxy.
	ProxyServerHost = "http.proxy.host"
	// ProxyServerPort is the port for the Kerberos reverse proxy.
	ProxyServerPort = "http.proxy.port"
	// APIServerHost is the host name for the Kerberos reverse proxy administration API.
	APIServerHost = "http.api.host"
	// APIServerPort is the port for the Kerberos reverse proxy administration API.
	APIServerPort = "http.api.port"
	// DatabaseServerHost is the host name for the Kerberos MySQL database server.
	DatabaseServerHost = "database.host"
	// DatabaseServerPort is the port for the the Kerberos MySQL database server.
	DatabaseServerPort = "database.port"
	// DatabaseServerUser is the username for the the Kerberos MySQL database server.
	DatabaseServerUser = "database.user"
	// DatabaseServerPass is the password for the the Kerberos MySQL database server.
	DatabaseServerPass = "database.pass"
	// DatabaseServerDBName is the name of the database for the the Kerberos MySQL database server.
	DatabaseServerDBName = "database.dbname"
	// PasswordHashCost is the cost of the password hashing algorithm (BCrypt)
	PasswordHashCost = "misc.password_hash_cost"
)

// ASCIILogo is the ascii representation of the Athena logo
var ASCIILogo = `
  _  __ ______  _____   ____   ______  _____    ____    _____ 
 | |/ /|  ____||  __ \ |  _ \ |  ____||  __ \  / __ \  / ____|
 | ' / | |__   | |__) || |_) || |__   | |__) || |  | || (___  
 |  <  |  __|  |  _  / |  _ < |  __|  |  _  / | |  | | \___ \ 
 | . \ | |____ | | \ \ | |_) || |____ | | \ \ | |__| | ____) |
 |_|\_\|______||_|  \_\|____/ |______||_|  \_\ \____/ |_____/ 

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
	setConfig(APIServerHost, "127.0.0.1")
	setConfig(APIServerPort, 8971)
	setConfig(DatabaseServerHost, "127.0.0.1")
	setConfig(DatabaseServerPort, 3306)
	setConfig(DatabaseServerUser, "root")
	setConfig(DatabaseServerPass, "root")
	setConfig(DatabaseServerDBName, "kerberos")
	setConfig(PasswordHashCost, 15)
}

// WriteConfig replaces the config file by the current configuration.
func WriteConfig() {
	filePath := viper.ConfigFileUsed()
	_ = os.Remove(filePath)
	CreateFile(filePath)
	err := viper.WriteConfig()
	if nil != err {
		log.Fatalln(err)
	}
}
