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
	// KerberosHost is the host name for the Kerberos reverse proxy.
	KerberosHost = "kerberos.host"
	// KerberosPort is the port for the Kerberos reverse proxy.
	KerberosPort = "kerberos.port"
	// KerberosAPIPort is the port for the Kerberos reverse proxy administration API.
	KerberosAPIPort = "kerberos.api_port"
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
	setConfig(KerberosHost, "127.0.0.1")
	setConfig(KerberosPort, 8970)
	setConfig(KerberosAPIPort, 8971)
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
