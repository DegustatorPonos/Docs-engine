package config

import (
	"log"
	"os"
	"strings"
)

var ConfigFileContents []string

//This part generates values from config.

func ReadConfigFile() {
	content, err := os.ReadFile("Config.cfg")
	if err != nil {
		log.Fatal(err)
	}
	ConfigFileContents = strings.Split(string(content), "\n")
}

// TODO: Rewrite it
// Returns a string that contains a given title. If there is none returns default value
func GetConfigParam(paramName string, defaultValue string) string {
	for _, cfgLine := range ConfigFileContents {
		if strings.Contains(cfgLine, paramName) {
			return strings.Replace(cfgLine, paramName+": ", "", -1)
		}
	}
	return defaultValue
}

// Returns port determined in config file. Returns default port in it is mot specified
func GetPort(defaultString string) string {
	return GetConfigParam("Port", defaultString)
}
