package resource

import "github.com/kelseyhightower/envconfig"

type LogConfig struct {
	Out    string `default:"file"`
	Level  string `default:"debug"`
	Format string `default:"json"`
	Path   string `default:"application.log"`
}

type Config struct {
	Server struct {
		Port int `default:"8080"`
	}

	DB struct {
		Host         string
		Port         uint `default:"3306"`
		User         string
		Password     string
		Name         string
		Charset      string `default:"utf8mb4"`
		Location     string `default:"UTC"`
		MaxIdleConns int    `default:"10"`
		MaxOpenConns int    `default:"10"`
		LogMode      bool   `default:"true"`
	}

	Log LogConfig
}

func LoadConfig() (*Config, error) {
	c := Config{}
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
