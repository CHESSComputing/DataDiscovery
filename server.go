package main

import (
	"fmt"
	"log"

	authz "github.com/CHESSComputing/golib/authz"
	srvConfig "github.com/CHESSComputing/golib/config"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// GET routes
	// all POST methods ahould be authorized
	authorized := r.Group("/")
	authorized.Use(authz.TokenMiddleware(srvConfig.Config.Authz.ClientID, srvConfig.Config.MetaData.Verbose))
	{
		authorized.GET("/", DataHandler)
	}

	return r
}

func Server() {
	// setup web router and start the service
	r := setupRouter()
	sport := fmt.Sprintf(":%d", srvConfig.Config.MetaData.WebServer.Port)
	log.Printf("Start HTTP server %s", sport)
	r.Run(sport)
}
