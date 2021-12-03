package config

import (
	gConfig "github.com/golobby/config/v3"
	"github.com/golobby/config/v3/pkg/feeder"
)

type config struct {
	Servers   serversConfig   `json:"servers"`
	Databases databasesConfig `json:"databases"`
}

var Instance = &config{}

// returns list of your wanted config feeders
func feeders() []gConfig.Feeder {
	return []gConfig.Feeder{
		&feeder.Json{Path: "app_config.json"},
	}
}
