
# Welcome to Marvel Characters API!    
 **Author: Dr Ahmad Hassan**      
 **Description** The Marvel Comics API to browse through the list of characters.      
 **Features** 
  - GET /character 
  - GET /character/{id}
  - Swagger UI
  - Docker Compose for building service and backing service i.e. redis
  - Test suite including controller tests
  
  **Caching Strategies**    
 I have implemented two caching strategies which are known as **TTL** and **PREFETCH**    
      
 1. **PREFETCH**     
 In PREFETCH caching strategy, the cache keys will have no cache expiry and cache will automatically be populated by pre-fetching with latest copy of the data after the configured time in config.yml. The default caching strategy is PREFETCH. The API can be started with PREFETCH as follows:    
    - go run main.go -s PREFETCH    
    
 2. **TTL**    
 In TTL caching strategy, the cache keys will automatically expire after the configured TTL in config.yml configuration. Subsequent API requests after the cache expiry will incur the cache miss and re-populate the cache.    
    - go run main.go -s TTL    
    
**Caching Assumptions** The assumption is that only the Marvel API is using the Redis cache exclusively so we can safely use FLUSHALL to delete all the keys and re-populate the cache by prefetching.    
   
 **How to RUN**  
 1. **FAST and EASY WAY TO START API using docker-compose**  
  
	  ```docker-compose up --build```  
	  
  2. **Manual way of running API** If you are still keen on running the API manually, then please follow the following steps to run the application    
    
   **Go version**    
     
 1.17 darwin/amd64      
         
   **Install routing package gorilla/mux** 
   
    go get -v -u github.com/gorilla/mux      
         
   **Install go-redis** 
   
   ```go get github.com/go-redis/redis ```     
         
   **Install cron**   
 ```
 go get github.com/robfig/cron/v3@v3.0.0  
go get github.com/robfig/cron/v3
``` 
  
  **Install swag**  
  ```  
 go get -u github.com/swaggo/swag/cmd/swag   
 go get github.com/swaggo/http-swagger      
   ```  
  **Run redis container**  
  ```  
 docker-compose up redis  
 ```
  **Run**   
  
    - go run main.go -s PREFETCH       
    - go run main.go -s TTL  
 

**Open browser**
 http://localhost:9091/swagger/index.html
 
 ** Tests **
docker-compose up redis 
go clean -testcache 
go test -v ./...

** Test Report **
PASS
ok  	github.com/marvel/controller	2.006s
=== RUN   TestHTTPError400
--- PASS: TestHTTPError400 (0.00s)
=== RUN   TestHTTPError404
--- PASS: TestHTTPError404 (0.00s)
=== RUN   TestHTTPError500
--- PASS: TestHTTPError500 (0.00s)
PASS
ok  	github.com/marvel/model	0.267s
=== RUN   TestGetMD5Hash
--- PASS: TestGetMD5Hash (0.00s)
=== RUN   TestGetCharacterIdUrl
--- PASS: TestGetCharacterIdUrl (0.00s)
PASS
ok  	github.com/marvel/utils	0.456s