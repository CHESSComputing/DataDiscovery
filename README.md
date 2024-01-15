# DataDiscovery

![build status](https://github.com/CHESSComputing/DataDiscovery/actions/workflows/go.yml/badge.svg)
[![go report card](https://goreportcard.com/badge/github.com/CHESSComputing/DataDiscovery)](https://goreportcard.com/report/github.com/CHESSComputing/DataDiscovery)
[![godoc](https://godoc.org/github.com/CHESSComputing/DataDiscovery?status.svg)](https://godoc.org/github.com/CHESSComputing/DataDiscovery)

CHESS Data Discovery service

### Example
```
# obtain valid token

# place search query request
curl -X POST \
    -H "Authorization: bearer $token" \
    -H "Content-type: application/json" \
    -d '{"client":"go-client","service_query":{"query":"{}","idx":0,"limit":2}}' \
    http://localhost:8320/search
```
