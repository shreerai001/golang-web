package routes

import (
	"sigma/api/controllers"

	"goji.io"
	"goji.io/pat"
)

func GetRegisteredRoutes() *goji.Mux {
	mux := getMux()
	mux.HandleFunc(pat.Get("/hello/:name"), controllers.HelloController)

	return mux
}
