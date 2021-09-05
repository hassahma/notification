package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/marvel/model"
	"fmt"
)

var PREFIX = "MARVEL_"

var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Exists(key string) bool {
	var ctx = context.Background()
	return rdb.Exists(ctx, PREFIX+key).Val() != 0
}

func DeleteAll() {
	var ctx = context.Background()
	rdb.FlushAll(ctx)
}

func Get(key string) model.Response {
	var ctx = context.Background()

	val, err := rdb.Get(ctx, PREFIX+key).Result()
	if err != nil {
		fmt.Println(err)
	}
	var res model.Response

	json.Unmarshal([]byte(val), &res)
	return res
}

func Set(key string, responseObject model.Response) interface{} {
	var ctx = context.Background()

	var err error

	b, err := json.Marshal(responseObject)
	err = rdb.Set(ctx, PREFIX+key, b, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
