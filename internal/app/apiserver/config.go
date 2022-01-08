package apiserver

import (
	"encoding/json"
	"io/ioutil"

	"github.com/serg-pe/blockchainrest/pkg/logging"
)

type ServerConfig struct {
	ServerAddress string `json:"serverAddress"`
	EthEndpoint   string `json:"ethEndpoint"`
}

func CreateConfigFromFile(filePath string) *ServerConfig {
	configBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		logging.GetLogger().Fatal("Can't read configs file")
	}

	var config ServerConfig
	err = json.Unmarshal(configBytes, &config)
	if err != nil {
		logging.GetLogger().Fatal("Can't unmarshal configs")
	}
	return &config
}
