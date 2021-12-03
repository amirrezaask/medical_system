package config

type serverConfig struct {
	Type string `json:"type"`
	Addr string `json:"addr"`
}
type serversConfig []serverConfig

func (s serversConfig) Get(typ string) serverConfig {
	for _, c := range s {
		if c.Type == typ {
			return c
		}
	}
	return serverConfig{}
}

type DatabaseConfig struct {
	DBName   string `json:"dbname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Host     string `json:"connection"`
	Port     int    `json:"port"`
	SSLMode  string `json:"sslmode"`
}
type databasesConfig map[string]DatabaseConfig
