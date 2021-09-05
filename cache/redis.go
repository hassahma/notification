package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/marvel/model"
	"github.com/marvel/constant"
	"github.com/marvel/utils"
	"fmt"
)


var rdb *redis.Client

func Init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     utils.Cfg.Redis.Url,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func Exists(key string) bool {
	var ctx = context.Background()
	return rdb.Exists(ctx, constant.PREFIX + key).Val() != 0
}

func DeleteAll() {
	var ctx = context.Background()
	rdb.FlushAll(ctx)
}

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

func Set(key string, responseObject model.Response) interface{} {
	var ctx = context.Background()

	var err error

	b, err := json.Marshal(responseObject)
	err = rdb.Set(ctx, constant.PREFIX + key, b, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}
