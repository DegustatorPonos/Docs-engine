package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var ConfigFileContents []string

//This part generates values from config.

func ReadConfigFile() {
	EnsureConfigFileCreated()
	content, err := os.ReadFile("Config.cfg")
	if err != nil {
		log.Fatal(err)
	}
	ConfigFileContents = strings.Split(string(content), "\n")
}

func EnsureConfigFileCreated() {
	if _, err := os.Stat("Config.cfg"); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("Recreating default config file")
		file, err := os.Create("Config.cfg")
		if err != nil {
			fmt.Println("Error while creating config file")
		}
		fmt.Fprint(file, "Port: 5000\nParsingMode: 0\nSource: ../Content/\nAppname: Paketik docs engine\n")
	}
}

// TODO: Rewrite it
// Returns a string that contains a given title. If there is none returns default value
func GetConfigParam(paramName string, defaultValue string) string {
	for _, cfgLine := range ConfigFileContents {
		if strings.Contains(cfgLine, paramName) {
			outp := strings.Replace(cfgLine, paramName+": ", "", -1)
			return strings.Trim(outp, " ")
		}
	}
	return defaultValue
}

// Returns port determined in config file. Returns default port in it is mot specified
func GetPort(defaultString string) string {
	return GetConfigParam("Port", defaultString)
}

// Returns port determined in config file. Returns default port in it is mot specified
func GetParsingMode(defaultMode int) int {
	mode, err := strconv.Atoi(GetConfigParam("ParsingMode", (string)(defaultMode)))
	if err != nil || mode > 2 {
		fmt.Println("Error while reading the mode value. The parsing mode is set to " + (string)(defaultMode))
		return defaultMode
	}
	return mode
}

// Returns path of the directory thet contains .md files
func GetSourceDirectoryPath(defaultPath string) string {
	return GetConfigParam("Source", defaultPath)
}

// Returns the name of the app
func GetAppName(defaultPath string) string {
	return GetConfigParam("Appname", defaultPath)
}
