package openstack_ssh

import (
	"github.com/BurntSushi/toml"
)

var configFilePath = "/etc/ssh/openstack-ssh.conf"

type Config struct {
	Auth_Url string
	User     string
	Password string
	Tenant   string
	Region   string
}

func LoadConfig() (*Config, error) {
	var config Config
	defaultConfig(&config)
	_, err := toml.DecodeFile(configFilePath, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func defaultConfig(config *Config) {
	config.Region = "regionOne"
}
