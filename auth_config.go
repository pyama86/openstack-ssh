package openstack_ssh

import "github.com/BurntSushi/toml"

const filePath = "/etc/openstack-ssh.conf"

type Config struct {
	Auth_Url string
	User     string
	Password string
	Tenant   string
	Region   string
}

func LoadConfig() *Config {
	var config Config
	defaultConfig(&config)
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		panic(err)
	}
	return &config
}

func defaultConfig(config *Config) {
	config.Region = "regionOne"
}
