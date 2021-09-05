
# Welcome to Marvel Characters API!    
 **Author: Dr Ahmad Hassan**      
 **Description** The Marvel Comics API to browse through the list of characters.      
      
  **Caching Strategies**    
 I have implemented two caching strategies which are known as **TTL** and **PREFETCH**    
      
 1. **PREFETCH**     
 In PREFETCH caching strategy, the cache keys will have no cache expiry and cache will automatically be populated by pre-fetching with latest copy of the data after the configured time in config.yml. The default caching strategy is PREFETCH. The API can be started with PREFETCH as follows:    
    - go run *.go -s PREFETCH    
    
 2. **TTL**    
 In TTL caching strategy, the cache keys will automatically expire after the configured TTL in config.yml configuration. Subsequent API requests after the cache expiry will incur the cache miss and re-populate the cache.    
    - go run *.go -s TTL    
    
**Caching Assumptions** The assumption is that only the Marvel API is using the Redis cache exclusively so we can safely use FLUSHALL to delete all the keys and re-populate the cache by prefetching.    
   
 **How to RUN**  
 1. **FAST and EASY WAY TO START API using docker-compose**  
  
	  ```docker-compose up --build```  
  2. **Manual way of running API** If you are still keen on running the API manually, then please follow the following steps to run the application    
    
   **Go version**    
     
 1.17 darwin/amd64      
         
   **Install routing package gorilla/mux** 
	   ```go get -v -u github.com/gorilla/mux```      
         
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
 docker-compose up redis    ```  
 ```
  **Run**   
  
    - go run *.go -s PREFETCH       
    - go run *.go -s TTL  
 

**Open browser**
 http://localhost:9091/swagger/index.html