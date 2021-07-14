package config

import (
	"fmt"
	"net/http"

	"sigma/api/routes"
)

func RunHttpServer(addr string) {
	mux := routes.GetRegisteredRoutes()

	fmt.Printf("************** Listening to port %s **************\n", string(addr))
	http.ListenAndServe("0.0.0.0:"+string(addr), mux)
}
