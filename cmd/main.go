package main

import (
	"flag"

	app "github.com/pyama86/openstack-ssh"
)

func main() {
	flag.Parse()
	config := app.LoadConfig()
	app.FetchPublicKey(flag.Arg(0), config)
}
