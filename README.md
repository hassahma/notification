
Author: Dr Ahmad Hassan

##The Marvel Comics API allows developers to access information on Marvel characters. In

##### go version
1.17 darwin/amd64

##### install routing package gorilla/mux
go get -v -u github.com/gorilla/mux

##### install go-redis
go get github.com/go-redis/redis

##### install cron
go get github.com/robfig/cron/v3@v3.0.0
go get github.com/robfig/cron/v3

##### install swag
go get -u github.com/swaggo/swag/cmd/swag
go get github.com/swaggo/http-swagger
swag init

##### run redis container
docker-compose up

##### run
go run *.go -s TTL
go run *.go -s PREFETCH

#### open browser
http://localhost:9091/swagger/index.html