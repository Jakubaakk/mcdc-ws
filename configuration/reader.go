package configuration

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

const DEFAULT_CONFIG_PATH = "./config.json"
const CONFIG_PATH_ENV_PROPERTY = "CONFIG_PATH"

func ReadConfiguration() McdcConfiguration {
	pathFromEnv := os.Getenv(CONFIG_PATH_ENV_PROPERTY)
	var path string
	if pathFromEnv == "" {
		path = DEFAULT_CONFIG_PATH
	} else {
		path = pathFromEnv
	}

	content, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Could not open configuration file: ", path, err)
	}

	var payload McdcConfiguration
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatalln("Could not unmarshall configuration file: ", path, err)
	}

	return payload
}
