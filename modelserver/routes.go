package modelserver

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strings"
)

type UserData struct {
	User string
	Data map[string]string
}

type Users struct {
	UsersList []UserData
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

func (c Config) RenderResponse(w http.ResponseWriter, r *http.Request, users Users) {
	js, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func (c Config) GetHandlerFunc() HandlerFunction {
	return func(w http.ResponseWriter, r *http.Request) {
		// todo: abstract pagedata provider interface
		path := strings.TrimRight(r.URL.Path, "/")
		switch path[1:] {
		case "users":
			users := make([]UserData, 1)
			users[0] = UserData{"alex", map[string]string{"name": "Aleksandr"}}
			c.RenderResponse(w, r, Users{users})
			break
		default:
			http.Error(w, errors.New("Not found").Error(), 404)
			break
		}

	}
}
