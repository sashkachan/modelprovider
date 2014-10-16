package modelserver

import (
	"fmt"
	"net/http"
)

type Config struct {
	PkgDir string
	Port   string
}

func StartServer(cfg Config) {
	fmt.Println("Starting server on :" + cfg.Port)
	routerProvider := RouterProvider(cfg)
	http.Handle("/", routerProvider.GetRouter())
	http.ListenAndServe("127.0.0.1:"+cfg.Port, nil)
}
