package config

import (
	"encoding/json"
	"io/ioutil"
)

type AutoRobotConfig struct {
	MeiRi    string
	Accounts []Account
}

type Account struct {
	AccountName string
	Password    string
	Roles       []Role
}

type Role struct {
	Id     int
	Fubens []FuBen
}

type FuBen struct {
	Name       string
	Difficulty int
}

func InitConfig(configPath string) (config *AutoRobotConfig, err error) {
	file, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	config = &AutoRobotConfig{}
	err = json.Unmarshal(file, config)
	return config, err
}
