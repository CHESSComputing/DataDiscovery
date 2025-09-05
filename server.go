package main

// server module
//
// Copyright (c) 2023 - Valentin Kuznetsov <vkuznet@gmail.com>
//
import (
	"log"

	srvConfig "github.com/CHESSComputing/golib/config"
	ql "github.com/CHESSComputing/golib/ql"
	server "github.com/CHESSComputing/golib/server"
	services "github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

var _httpReadRequest *services.HttpRequest
var _qlMgr ql.QLManager
var Verbose int

// helper function to setup our router
func setupRouter() *gin.Engine {
	routes := []server.Route{
		{Method: "GET", Path: "/", Handler: DataHandler, Authorized: true},
		{Method: "POST", Path: "/nrecords", Handler: NRecordsHandler, Authorized: true},
		{Method: "POST", Path: "/search", Handler: SearchHandler, Authorized: true},
	}
	r := server.Router(routes, nil, "static", srvConfig.Config.Discovery.WebServer)
	return r
}

// Server defines our HTTP server
func Server() {
	Verbose = srvConfig.Config.Discovery.WebServer.Verbose
	_httpReadRequest = services.NewHttpRequest("read", Verbose)

	// initialize QL manager
	if err := _qlMgr.Init(srvConfig.Config.QL.ServiceMapFile); err != nil {
		log.Println("ERROR: FOXDEN QL Manager is not initialized", err)
	}

	// setup web router and start the service
	r := setupRouter()
	webServer := srvConfig.Config.Discovery.WebServer
	server.StartServer(r, webServer)
}
