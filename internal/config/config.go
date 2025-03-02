package config

import (
	"encoding/json"
	"os"

	"github.com/fermar/gator/internal/logging"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFile string = ".gatorconfig.json"

func Read() (*Config, error) {
	var localDir string
	localDir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	localDir = localDir + "/" + configFile
	logging.Lg.Logger.Printf("Leyendo archivo local: %v ...\n", localDir)
	rawConf, err := os.ReadFile(localDir)
	if err != nil {
		return nil, err
	}
	conf := Config{}
	err = json.Unmarshal(rawConf, &conf)
	if err != nil {
		return nil, err
	}
	// logging.Lg.Logger.Printf("configuracion: %v\n", conf)
	return &conf, nil
}

func (conf *Config) SetUser(username string) error {
	conf.CurrentUserName = username
	jEncoded, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	var localFile string
	localFile, err = os.UserHomeDir()
	if err != nil {
		return err
	}

	localFile = localFile + "/" + configFile
	// logging.Lg.Logger.Printf("Leyendo archivo local: %v ...\n", localFile)
	err = os.WriteFile(localFile, jEncoded, 0644)
	if err != nil {
		return err
	}
	return nil
}
