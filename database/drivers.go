package database

import (
	"database/sql"
	"medical_system/config"
)

var Drivers = map[string]*sql.DB{}

func connect(cfg config.DatabaseConfig) (*sql.DB, error) {
	return nil, nil
}

func init() {
	cfg := config.Instance
	for name, dbCfg := range cfg.Databases {
		db, err := connect(dbCfg)
		if err != nil {
			panic(err)
		}
		Drivers[name] = db
	}
}
