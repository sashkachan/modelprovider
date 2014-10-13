package modelprovider

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type UserData struct {
	User string
	Data map[string]string
}

type HandlerFunction func(http.ResponseWriter, *http.Request)

type RouterProvider interface {
	GetRouter() *mux.Router
	GetHandlerFunc() HandlerFunction
}

func (c Config) GetRouter() *mux.Router {
	fmt.Println("Getting routes")
	router := mux.NewRouter()
	router.HandleFunc("/{action}", c.GetHandlerFunc())
	return router
}

func (c Config) RenderResponse(w http.ResponseWriter, r *http.Request, user UserData) {
	js, err := json.Marshal(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
	// path := r.URL.Path
	// trimmedPath := strings.TrimRight(path, "/")

	// templatePath, ok := c.TemplatesMap[trimmedPath]

	// if ok == false {
	// 	fmt.Println("Throwing template error")
	// 	return nil, errors.New("Template not found")
	// }
	// assetsDir := c.PkgDir + "/" + "assets" + "/"
	// fmt.Println(assetsDir + templatePath)
	// tmpl, err := template.ParseFiles(assetsDir + templatePath)

	// if err != nil {
	// 	fmt.Println("Going to explode now!!")
	// 	http.Redirect(w, r, "/", http.StatusFound)
	// }
	// if err = tmpl.Execute(w, pdata); err != nil {
	// 	panic(err)
	// }
}

func (c Config) GetHandlerFunc() HandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo: abstract pagedata provider interface
		pdata := UserData{"alex", map[string]string{"name": "Aleksandr"}}
		c.RenderResponse(w, r, pdata)
	}
}
