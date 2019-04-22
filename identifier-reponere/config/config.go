package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	MongoUrl   string
	Username   string
	Password   string
	Name       string
	Collection string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			MongoUrl:   "192.168.99.103:27017",
			Username:   "admin",
			Password:   "admin",
			Name:       "smart-mongo",
			Collection: "object-repository",
		},
	}
}
