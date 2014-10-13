package modelprovider

import (
	"fmt"
	"net/http"
)

type Config struct {
	PkgDir       string
	TemplatesMap map[string]string
}

func StartServer(cfg Config) {
	fmt.Println("Starting server on :8080")
	routerProvider := RouterProvider(cfg)
	http.Handle("/", routerProvider.GetRouter())
	http.ListenAndServe(":8080", nil)
}
