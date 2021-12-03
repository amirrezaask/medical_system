package config

import (
	gConfig "github.com/golobby/config/v3"
)

func getInstance() *config {
	fed := &config{}
	c := gConfig.New()
	for _, feeder := range feeders() {
		c.AddFeeder(feeder)
	}
	c.AddStruct(fed)
	c.Feed()
	return fed
}

func init() {
	Instance = getInstance()
}
