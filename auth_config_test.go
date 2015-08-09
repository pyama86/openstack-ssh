package openstack_ssh

import (
	"bytes"
	"os"
	"testing"

	"github.com/BurntSushi/toml"
)

func TestLoadConfig(t *testing.T) {
	configFilePath = "/tmp/openstack-ssh.conf"
	_config := Config{User: "pyama", Password: "testpass", Auth_Url: "http://testurl", Tenant: "testtenant", Region: "testregion"}

	var buffer bytes.Buffer
	encoder := toml.NewEncoder(&buffer)
	err := encoder.Encode(_config)
	if err != nil {
		panic(err)
	}

	f, err := os.Create(configFilePath)
	if _, err := buffer.WriteTo(f); err != nil {
		t.Fatalf("UnCreated File:%s", configFilePath)
	}

	config, err := LoadConfig()
	if err != nil {
		t.Fatalf("UnLoaded File:%s", configFilePath)
	}

	if config.User != "pyama" {
		t.Fatalf("Loading user should be failed")
	}

	if config.Password != "testpass" {
		t.Fatalf("Loading password should be failed")
	}

	if config.Auth_Url != "http://testurl" {
		t.Fatalf("Loading auth_url should be failed")
	}

	if config.Tenant != "testtenant" {
		t.Fatalf("Loading tenant should be failed")
	}

	if config.Region != "testregion" {
		t.Fatalf("Loading region should be failed")
	}
}
