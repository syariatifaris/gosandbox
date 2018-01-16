package main

import (
	"log"
	"net/http"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gorilla/mux"
	"github.com/syariatifaris/gosandbox/handler"
	"github.com/syariatifaris/gosandbox/inject"
)

func main() {
	dependencies := inject.NewDependencies()

	var controllers []handler.THandler
	var muxRouter *mux.Router

	inject.GetAssignedDependencies(dependencies, &controllers)
	inject.GetAssignedDependency(dependencies, &muxRouter)

	for _, c := range controllers {
		log.Println("Registering controller:", c.Name())
		c.RegisterHandlers(muxRouter)
	}

	// serve and listen for shutdown signals
	gracehttp.Serve(
		&http.Server{Addr: "0.0.0.0:8080", Handler: muxRouter},
	)
}
