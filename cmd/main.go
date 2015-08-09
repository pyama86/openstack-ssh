package main

import (
	"flag"
	"fmt"

	app "github.com/pyama86/openstack-ssh"
)

func main() {
	flag.Parse()
	config := app.LoadConfig()
	key, err := app.FetchPublicKey(flag.Arg(0), config)

	if err != nil {
		panic(err)
	}

	fmt.Println(key.PublicKey)
}
