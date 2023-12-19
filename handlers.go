package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	srvConfig "github.com/CHESSComputing/golib/config"
	services "github.com/CHESSComputing/golib/services"
	"github.com/gin-gonic/gin"
)

// DataHandler provives access to GET / end-point
func DataHandler(c *gin.Context) {
	rec := services.Response("DataDiscovery", http.StatusOK, services.OK, nil)
	c.JSON(200, rec)
}

// SearchHandler provives access to GET / end-point
func SearchHandler(c *gin.Context) {
	r := c.Request
	defer r.Body.Close()
	data, err := io.ReadAll(r.Body)
	if err != nil {
		rec := services.Response("DataDiscovery", http.StatusBadRequest, services.ReaderError, err)
		c.JSON(http.StatusBadRequest, rec)
		return
	}

	// TODO: implement logic to search across different services like
	// MetaData, DataBookkeeping, ScanService, etc.
	// so far we query MetaData service
	rurl := fmt.Sprintf("%s/search", srvConfig.Config.Services.MetaDataURL)
	resp, err := _httpReadRequest.Post(rurl, "application/json", bytes.NewBuffer(data))
	if err != nil {
		rec := services.Response("DataDiscovery", http.StatusBadRequest, services.ServiceError, err)
		c.JSON(http.StatusBadRequest, rec)
		return
	}
	c.JSON(200, resp)
}
