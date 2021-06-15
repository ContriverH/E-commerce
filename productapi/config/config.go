package config

import (
	"os"
)

// Config is the server configuration structure.
// all fields will be filled with environment variables.
type Config struct {
	ServerHost    string // address that server will listening on
	MongoUser     string // mongo db username
	MongoPassword string // mongo db password
	MongoHost     string // host that mongo db listening on
	MongoPort     string // port that mongo db listening on
}

// initialize will read environment variables and save them in config structure fields
func (config *Config) initialize() {
	// read environment variables
	// config.ServerHost = os.Getenv("http://localhost:8081")
	config.MongoUser = os.Getenv("mongo_user")
	config.MongoPassword = os.Getenv("mongo_password")
	config.MongoHost = os.Getenv("mongo_host")
	config.MongoPort = os.Getenv("mongo_port")
}

// NewConfig will create and initialize config struct
func NewConfig() *Config {
	config := new(Config)
	config.initialize()
	return config
}
