package main

import (
	"fmt"
	"flag"
)

func main() {
	flag.Parse()
        if flag.Arg(0) == "pyama" {
		fmt.Println("ssh hogehoge")
	}
}
