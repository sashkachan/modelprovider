package modelprovider

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strings"
)

type PageData struct {
	Title, Body string
}

type HandlerFunction func(http.ResponseWriter, *http.Request)

type RouterProvider interface {
	GetRouter() *mux.Router
	GetHandlerFunc() HandlerFunction
}

func (c Config) GetRouter() *mux.Router {
	fmt.Println("Getting routes")
	router := mux.NewRouter()
	router.HandleFunc("/", c.GetHandlerFunc())
	router.HandleFunc("/{action}", c.GetHandlerFunc())
	return router
}

func (c Config) RenderResponse(request *http.Request) (*template.Template, error) {
	path := request.URL.Path
	trimmedPath := strings.TrimRight(path, "/")
	templatePath, ok := c.TemplatesMap[trimmedPath]
	if ok == false {
		panic("Template not found")
	}
	assetsDir := c.PkgDir + "/" + "assets" + "/"
	return template.ParseFiles(assetsDir + templatePath)
}

func (c Config) GetHandlerFunc() HandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo: abstract pagedata provider interface
		pdata := PageData{}
		tmpl, err := c.RenderResponse(r)
		if err != nil {
			panic(err)
		}
		if err = tmpl.Execute(w, pdata); err != nil {
			panic(err)
		}
	}
}
