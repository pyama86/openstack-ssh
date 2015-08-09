package openstack_ssh

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/compute/v2/extensions/keypairs"
)

type KeyPair struct {
	keypairs.GetResult
}

func FindKeyPairByName(cli *gophercloud.ServiceClient, name string) (*KeyPair, error) {
	keypair := keypairs.Get(cli, name)
	return &KeyPair{keypair}, nil
}
