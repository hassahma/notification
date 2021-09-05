package cache

import (
"time"
	"encoding/json"
r "gopkg.in/redis.v5"
)

var preffix = "MARVEL_CHARACTERS_"

type Storage struct {
	client *r.Client
}

//NewStorage creates a new redis storage
func NewStorage(url string) (*Storage, error) {
	var (
		opts *r.Options
		err  error
	)

	if opts, err = r.ParseURL(url); err != nil {
		return nil, err
	}

	return &Storage{
		client: r.NewClient(opts),
	}, nil
}

//Get a cached content by key
func (s Storage) Get(key string) []byte {
	val, _ := s.client.Get(preffix + key).Bytes()
	return val
}

//Set a cached content by key
func (s Storage) Set(key string, value interface{}, duration time.Duration) {
	p, err := json.Marshal(value)
	if err != nil {
	}
	s.client.Set(preffix+key, p, duration)
}
