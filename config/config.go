package config

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	Host     string   `json:"host"`
	Postgres Postgres `json:"postgres"`
}

func ReadConfig(filename *string) (conf Config, err error) {
	confData, err := ioutil.ReadFile(*filename)
	if err != nil {
		return conf, err
	}
	if err = json.Unmarshal(confData, &conf); err != nil {
		return conf, err
	}
	return conf, nil
}

type Postgres struct {
	Dialect      string `json:"dialect"`
	User         string `json:"user"`
	DataBaseName string `json:"db_name"`
	SSLMode      string `json:"ssl_mode"`
	Password     string `json:"password"`
}
