package blog

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

func (c Config) GetTemplateByPath(request *http.Request) (*template.Template, error) {
	path := request.URL.Path
	trimmedPath := strings.TrimRight(path, "/")
	templatePath, ok := c.TemplatesMap[trimmedPath]
	if ok == false {
		templatePath = "index.html"
	}
	return template.ParseFiles(c.PkgDir + "/" + "assets" + "/" + templatePath)
}

func (c Config) GetHandlerFunc() HandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		pdata := PageData{}
		tmpl, err := c.GetTemplateByPath(r)
		if err != nil {
			panic(err)
		}
		if err = tmpl.Execute(w, pdata); err != nil {
			panic(err)
		}
	}
}
