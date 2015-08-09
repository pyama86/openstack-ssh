package main

import (
	"flag"
	"fmt"

	app "github.com/pyama86/openstack-ssh"
)

func main() {

	flag.Parse()
	if flag.Arg(0) == "" {
		panic("please input username")
	}
	config, err := app.LoadConfig()

	if err != nil {
		app.ERR(err)
		panic(err)
	}

	key, err := app.FetchPublicKey(flag.Arg(0), config)

	if err != nil {
		app.ERR(err)
		panic(err)
	}

	fmt.Println(key.PublicKey)
}
