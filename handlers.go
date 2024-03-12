package main

// handlers module
//
// Copyright (c) 2023 - Valentin Kuznetsov <vkuznet@gmail.com>
//
import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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
	// keep query data for multiple search requests
	query := data

	// to proceed obtain valid token
	_httpReadRequest.GetToken()

	// TODO: implement logic to search across different services like
	// MetaData, DataBookkeeping, ScanService, etc.
	// so far we query MetaData service

	// get number of total records for our query
	rurl := fmt.Sprintf("%s/count", srvConfig.Config.Services.MetaDataURL)
	resp, err := _httpReadRequest.Post(rurl, "application/json", bytes.NewBuffer(query))
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
	var nrecords int
	err = json.Unmarshal(data, &nrecords)
	if err != nil {
		log.Printf("provided data: %v failed to unmarshal to int data-type\n", string(data))
		rec := services.Response("DataDiscovery", http.StatusBadRequest, services.UnmarshalError, err)
		c.JSON(http.StatusBadRequest, rec)
		return
	}
	log.Println("### MetaData count:", nrecords, string(data))

	if nrecords == 0 {
		rec := services.ServiceResponse{
			Service:   "MetaData",
			Results:   &services.ServiceResults{NRecords: nrecords},
			Timestamp: time.Now().String(),
		}
		c.JSON(200, rec)
		return
	}

	// get results records
	rurl = fmt.Sprintf("%s/search", srvConfig.Config.Services.MetaDataURL)
	resp, err = _httpReadRequest.Post(rurl, "application/json", bytes.NewBuffer(query))
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

	// each service response contains only list of records or error record
	// therefore we'll wrap it in service response record
	var records []map[string]any
	err = json.Unmarshal(data, &records)
	if err != nil {
		log.Println("ERROR: response", string(data))
		rec := services.Response("DataDiscovery", http.StatusBadRequest, services.UnmarshalError, err)
		c.JSON(http.StatusBadRequest, rec)
		return
	}
	rec := services.ServiceResponse{
		Service:   "MetaData",
		Results:   &services.ServiceResults{NRecords: nrecords, Records: records},
		Timestamp: time.Now().String(),
	}
	c.JSON(200, rec)
}
