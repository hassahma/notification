// model contains all the structs that are used within the application
package model

import (
	"time"
)

// HTTPError400 for error 400
type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPError404 for error 404
type HTTPError404 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not Found"`
}

// HTTPError500 for error 500
type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}

// Character struct is returned for /characters/{id} api response
type Character struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

// Characters struct with id property only.
type Characters struct {
	Id int `json:"id"`
}

// Response is the struct used for de-serializing response from Marvel API
type Response struct {
	Code int    `json:"code"`
	Status string  `json:"status"`
	Data Data `json:"data"`
}

// struct for holding the 'data' part of json response
type Data struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Count int `json:"count"`
	Results []Character `json:"results"`
}

// Config struct for reading the configuration file config.yml
type Config struct {
	Redis struct {
		Url string `yaml:"url"`
	} `yaml:"redis"`
	Marvel struct {
		Timestamp string `yaml:"timestamp"`
		Apikey string `yaml:"apikey"`
		Privatekey string `yaml:"privatekey"`
		Limit string `yaml:"limit"`
		TTL time.Duration `yaml:"TTL"`
		PREFETCH int `yaml:"PREFETCH"`
	} `yaml:"marvel"`
}
