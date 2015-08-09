package openstack_ssh

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
)

func FindKeyPairByName(cli *gophercloud.ServiceClient, name string) (*keypairs.KeyPair, error) {
	result := keypairs.Get(cli, name)
	keypair, err := result.Extract()
	if err != nil {
		panic(err)
	}
	return keypair, nil
}
