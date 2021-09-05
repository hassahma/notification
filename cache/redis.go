// cache package that implements a set of methods for storing and getting data from redis
package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/marvel/model"
	"github.com/marvel/constant"
	"github.com/marvel/utils"
	"fmt"
	"os"
)
//  Redis client
var rdb *redis.Client

// Initialize redis connection
func Init() {
	redis_url := os.Getenv("REDIS_DOCKER")
	if len(redis_url) == 0  {
		redis_url = utils.Cfg.Redis.Url
	}
	fmt.Println("\n Using redis: ", redis_url)
	rdb = redis.NewClient(&redis.Options{
		Addr:     redis_url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// Check if key exists in redis and returns bool.
func Exists(key string) bool {
	var ctx = context.Background()
	return rdb.Exists(ctx, constant.PREFIX + key).Val() != 0
}

// Delete all the keys from redis.
func DeleteAll() {
	var ctx = context.Background()
	rdb.FlushAll(ctx)
}

// Gets the keys
func Get(key string) model.Response {
	var ctx = context.Background()

	val, err := rdb.Get(ctx, constant.PREFIX + key).Result()
	if err != nil {
		fmt.Println(err)
	}
	var res model.Response

	json.Unmarshal([]byte(val), &res)
	return res
}

// Sets the key in redis with a chosen TTL strategy.
func Set(key string, responseObject model.Response) interface{} {
	var ctx = context.Background()

	var err error
	b, err := json.Marshal(responseObject)

	// Sets they key and gets the TTL based on the caching strategy used.
	err = rdb.Set(ctx, constant.PREFIX + key, b, utils.GetExpirationBasedOnStrategy()).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
