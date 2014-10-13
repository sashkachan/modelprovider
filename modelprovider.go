package main

import (
	"github.com/alex-glv/modelprovider/modelserver"
)

func main() {
	cfg := modelserver.Config{
		PkgDir: "/Users/alexg/go/src/github.com/alex-glv/modelprovider/modelserver/",
		TemplatesMap: map[string]string{
			"/users": "users.json",
		},
		Port: "8080",
	}
	modelserver.StartServer(cfg)
}
