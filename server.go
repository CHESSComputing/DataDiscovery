package main

import (
	"fmt"
	"log"

	authz "github.com/CHESSComputing/golib/authz"
	srvConfig "github.com/CHESSComputing/golib/config"
	services "github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

var _httpReadRequest *services.HttpRequest
var Verbose int

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// GET routes
	// all POST methods ahould be authorized
	authorized := r.Group("/")
	authorized.Use(authz.TokenMiddleware(srvConfig.Config.Authz.ClientID, srvConfig.Config.Discovery.Verbose))
	{
		authorized.GET("/", DataHandler)
		authorized.POST("/search", SearchHandler)
	}

	return r
}

func Server() {
	Verbose = srvConfig.Config.Discovery.WebServer.Verbose
	_httpReadRequest = services.NewHttpRequest("read", Verbose)
	// setup web router and start the service
	r := setupRouter()
	sport := fmt.Sprintf(":%d", srvConfig.Config.Discovery.WebServer.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}
