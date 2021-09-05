package main

import "time"

type CacheInterface interface {
	Get(key string) []byte
	Set(key string, content interface{}, duration time.Duration)
}