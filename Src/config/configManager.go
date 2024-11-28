package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var ConfigFileContents []string
var ConfigFileMap map[string]string = make(map[string]string)

//This part generates values from config.

func ReadConfigFile() error {
	EnsureConfigFileCreated()
	content, err := os.ReadFile("Config.cfg")
	if err != nil {
		log.Fatal(err)
	}
	ConfigFileContents = strings.Split(string(content), "\n")
	for _, configLine := range ConfigFileContents {
		var contents = strings.Split(configLine, ": ")
		if len(contents) != 2 {
			// Occurs when there are more ': ' separators in the file tan expected
			// TODO: Might be wrong here, so possbly a fix will be needed
			continue
		}
		// fmt.Println(contents)
		var paramName = strings.TrimSuffix(contents[0], " ")
		paramName = strings.TrimSuffix(paramName, ":")
		var paramValue = TrimEndSymbols(contents[1])
		ConfigFileMap[paramName] = paramValue
	}
	DumpMapValueInConsole(ConfigFileMap)
	return nil
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

// Returns a string that contains a given title. If there is none returns default value
func GetConfigParam(paramName string, defaultValue string) string {
	value, paramExists := ConfigFileMap[paramName]
	if !paramExists {
		fmt.Printf("Parameter does not exist in the nonfig file. Param name: %s\n", paramName)
		return defaultValue
	}
	return value
}

// Returns port determined in config file. Returns default port in it is mot specified
func GetPort(defaultString string) string {
	fmt.Println(GetConfigParam("Port", defaultString))
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

// Prints the contents of a map (including its keys and values) in the console
func DumpMapValueInConsole(inp map[string]string) {
	fmt.Println("map value dump:")
	for val := range inp {
		fmt.Printf("name \"%s\" ", val)
		kvp, err := ConfigFileMap[val]
		fmt.Printf("exists \"%v\" ", err)
		fmt.Printf("val \"%v\" end", kvp)
		fmt.Println()
	}
	fmt.Println("map value dump end")
}

// Trims unicode strings from the end of the line. Somehow this is the thing we have to deal with
func TrimEndSymbols(inp string) string {
	var runes = []rune(inp)
	var index int = len(runes) - 1
	for {
		var r = runes[index]
		if unicode.IsSpace(r) {
			runes[index] = 0
		} else {
			break
		}
		index -= 1
	}
	return string(runes)
}
