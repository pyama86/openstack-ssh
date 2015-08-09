package openstack_ssh

import (
	"fmt"

	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
)

func FetchPublicKey(userName string, config *Config) {

	fmt.Println(config.Region)
	regionOpts := gophercloud.EndpointOpts{Region: config.Region}

	provider := initClient(config)

	computeClient, err := openstack.NewComputeV2(provider, regionOpts)
	if err != nil {
		panic(err)
	}
	FindKeyPairByName(computeClient, userName)
}

func initClient(config *Config) *gophercloud.ProviderClient {
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: config.Auth_Url,
		Username:         config.User,
		Password:         config.Password,
		TenantName:       config.Tenant,
	}
	client, err := openstack.AuthenticatedClient(opts)
	if err != nil {
		panic(err)
	}
	return client
}
