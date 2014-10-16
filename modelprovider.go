package main

import (
	"flag"
	"github.com/alex-glv/modelprovider/modelserver"
)

func main() {
	portPrm := flag.String("port", "8912", "Port number to run http sever")
	flag.Parse() //

	cfg := modelserver.Config{
		PkgDir: "/Users/alexg/go/src/github.com/alex-glv/modelprovider/",
		Port:   *portPrm,
	}
	modelserver.StartServer(cfg)
}
