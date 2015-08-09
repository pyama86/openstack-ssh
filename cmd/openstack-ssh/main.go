package main

import (
	"flag"
	"fmt"

	app "github.com/pyama86/openstack-ssh"
)

const Version string = "0.1.0"

func main() {
	var versionFlg bool
	flag.BoolVar(&versionFlg, "v", false, "show version")
	flag.Parse()

	if versionFlg {
		fmt.Println("openstack-ssh version:", Version)
		return
	}
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
