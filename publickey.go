package openstack_ssh

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
)

func FetchPublicKey(userName string, config *Config) (*keypairs.KeyPair, error) {

	regionOpts := gophercloud.EndpointOpts{Region: config.Region}
	provider := initClient(config)
	computeClient, err := openstack.NewComputeV2(provider, regionOpts)

	if err != nil {
		panic(err)
	}

	key, err := FindKeyPairByName(computeClient, userName)

	if err != nil {
		panic(err)
	}

	return key, nil
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
