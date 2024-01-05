package main

import (
	"fmt"
	"log"

	srvConfig "github.com/CHESSComputing/golib/config"
	server "github.com/CHESSComputing/golib/server"
	services "github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

var _httpReadRequest *services.HttpRequest
var Verbose int

// helper function to setup our router
func setupRouter() *gin.Engine {
	routes := []server.Route{
		server.Route{Method: "GET", Path: "/", Handler: DataHandler, Authorized: true},
		server.Route{Method: "POST", Path: "/search", Handler: SearchHandler, Authorized: true},
	}
	r := server.Router(routes, nil, "static", srvConfig.Config.Discovery.WebServer)
	return r
}

// Server defines our HTTP server
func Server() {
	Verbose = srvConfig.Config.Discovery.WebServer.Verbose
	_httpReadRequest = services.NewHttpRequest("read", Verbose)
	// setup web router and start the service
	r := setupRouter()
	sport := fmt.Sprintf(":%d", srvConfig.Config.Discovery.WebServer.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}
