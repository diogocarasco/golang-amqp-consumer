package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var config *configuration

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

type appConfig struct {
	Port string
}

type configuration struct {
	APP  appConfig
	AMQP Config
}

// Returns the connection string
func (d *Config) GetConnectionString() string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s", d.User, d.Password, d.Host, d.Port)
}

func init() {

	viper.SetDefault("api.port", "8000")
	viper.SetDefault("db.host", "localhost")
	viper.SetDefault("db.port", "5432")

}

// Loads the configuration from config file
func Load() error {

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath("../config/")
	err := viper.ReadInConfig()
	if err != nil { // Handle errors reading the config file
		return fmt.Errorf("fatal error config file: %w", err)
	}

	config = new(configuration)

	return nil

}
