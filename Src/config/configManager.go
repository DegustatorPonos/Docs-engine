package config

import (
	"log"
	"os"
)

//This part generates values from config.

func ReadConfigFile() []byte {
	content, err := os.ReadFile("Config.cfg")
	if err != nil {
		log.Fatal(err)
	}
	return content
}
