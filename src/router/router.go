package router

import (
	"devbook-api/src/router/routes"
	"github.com/gorilla/mux"
)

func Gerar() *mux.Router {
	r := mux.NewRouter()
	return routes.Setup(r)
}
