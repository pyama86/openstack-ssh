package openstack_ssh

import (
	"fmt"
	"net/http"
	"testing"

	fake "github.com/rackspace/gophercloud/openstack/networking/v2/common"
	th "github.com/rackspace/gophercloud/testhelper"
)

const GetOutput = `
{
	"keypair": {
		"public_key": "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAAAgQC+Eo/RZRngaGTkFs7I62ZjsIlO79KklKbMXi8F+KITD4bVQHHn+kV+4gRgkgCRbdoDqoGfpaDFs877DYX9n4z6FrAIZ4PES8TNKhatifpn9NdQYWA+IkU8CuvlEKGuFpKRi/k7JLos/gHi2hy7QUwgtRvcefvD/vgQZOVw/mGR9Q== Generated by Nova\n",
		"name": "firstkey",
		"fingerprint": "15:b0:f8:b3:f9:48:63:71:cf:7b:5b:38:6d:44:2d:4a"
	}
}
`

func TestFindFlavorByName(t *testing.T) {
	th.SetupHTTP()
	defer th.TeardownHTTP()

	th.Mux.HandleFunc("/v2.0/os-keypairs/firstkey", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		r.ParseForm()
		fmt.Fprintf(w, GetOutput)
	})

	cli := fake.ServiceClient()

	keypair, err := FindKeyPairByName(cli, "firstkey")
	if err != nil {
		t.Fatalf("Unexpected error: %s", err.Error())
	}
	if keypair.Name != "firstkey" {
		t.Fatalf("Unmatch Name: %s", keypair.Name)
	}
}