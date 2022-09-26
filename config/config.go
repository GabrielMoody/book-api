package config

type Config struct {
	DB *DBConfig
}

type DBConfig struct {
	Dialect  string
	Host     string
	Port     int
	Username string
	Password string
	Name     string
	Charset  string
}

func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Dialect:  "mysql",
			Host:     "localhost",
			Port:     3306,
			Username: "root",
			Password: "",
			Name:     "books",
			Charset:  "utf8",
		},
	}
}
