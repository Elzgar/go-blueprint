package config

import "github.com/tkanos/gonfig"

// Configuration struct is for store data from config
type Configuration struct {
	DbUsername string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
}

// GetConfig function is for getting config data config.json and store at Configuration
func GetConfig() Configuration {
	conf := Configuration{}

	gonfig.GetConf("config/config.json", &conf)

	return conf
}
