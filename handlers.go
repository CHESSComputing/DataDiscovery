package main

// handlers module
//
// Copyright (c) 2023 - Valentin Kuznetsov <vkuznet@gmail.com>
//
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

	// to proceed obtain valid token
	_httpReadRequest.GetToken()

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

	// read respnose from downstream service
	defer resp.Body.Close()
	data, err = io.ReadAll(resp.Body)
	if err != nil {
		rec := services.Response("DataDiscovery", http.StatusBadRequest, services.ReaderError, err)
		c.JSON(http.StatusBadRequest, rec)
		return
	}

	/*
		var rec services.ServiceResponse
		err = json.Unmarshal(data, &rec)
		if err != nil {
			rec := services.Response("DataDiscovery", http.StatusBadRequest, services.UnmarshalError, err)
			c.JSON(http.StatusBadRequest, rec)
			return
		}
		c.JSON(200, rec)
	*/

	// extract content type from response header, if it is missing set default value
	ctype := resp.Header.Get("Content-type")
	if ctype == "" {
		ctype = "application/json"
	}
	c.Data(200, ctype, data)

	/*
		// TODO: I should replace reading response body from downstream service by
		// using reader closer

		tee := io.TeeReader(resp.Body, c.Writer)
		_, err = io.Copy(c.Writer, tee)
		if err != nil {
			rec := services.Response("DataDiscovery", http.StatusBadRequest, services.ServiceError, err)
			c.JSON(http.StatusBadRequest, rec)
		}
		c.JSON(200, nil)
	*/
}
