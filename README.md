# DataDiscovery
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
