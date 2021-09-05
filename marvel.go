package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
)

var marvel_url = "https://gateway.marvel.com/v1/public/characters"

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

var storage CacheInterface

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func buildURL(urlStr string, offset string) string {
	ts := "1"
	apikey := "17c400eff8dafac0184fb02420750089"
	privateKey := "1d724a31e223d1fd04a442b129c66b4b6528b360"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}

	url.Scheme = "https"
	queryparams := url.Query()
	queryparams.Set("ts", ts)
	queryparams.Set("apikey", apikey)
	queryparams.Set("hash", getMD5Hash(ts+privateKey+apikey))
	queryparams.Set("limit", "100")
	queryparams.Set("offset", offset)

	url.RawQuery = queryparams.Encode()
	fmt.Println("query is " + url.String())
	return url.String()
}
type Character struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type Characters struct {
	Id int `json:"id"`
}
var characters [] Character


type Response struct {
	Code int    `json:"code"`
	Status string  `json:"status"`
	Data Data `json:"data"`

}

type Data struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Count int `json:"count"`
	Results []Character `json:"results"`
}

type Output struct {
	Offset int `json:"offset"`
	Limit int `json:"limit"`
	Total int `json:"total"`
	Count int `json:"count"`
	Results []Characters `json:"results"`
}

// HTTPError
type HTTPError400 struct {
	Code    int    `json:"code" example:"400"`
	Message string `json:"message" example:"status bad request"`
}

// HTTPError
type HTTPError404 struct {
	Code    int    `json:"code" example:"404"`
	Message string `json:"message" example:"Not Found"`
}

// HTTPError
type HTTPError500 struct {
	Code    int    `json:"code" example:"500"`
	Message string `json:"message" example:"Internal Server Error"`
}

// GetAllCharacters godoc
// @Summary Serves an endpoint /characters that returns all the Marvel character ids in a JSON array of numbers.
// @Description Returns the IDs of all the Marvel Characters.
// @Tags Characters
// @Accept  json
// @Produce  json
// @Success 200 {array} int
// @Failure 400 {object} HTTPError400
// @Failure 500 {object} HTTPError500
// @Router /characters [get]
func getAllCharacters (w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all characters")
	json.NewEncoder(w).Encode(fetch())
}

func invalidateAndRefresh(){
	fmt.Println("invalidating cache")
	deleteAll()
	fmt.Println("prefetch cache")
	fetch()
}
func fetch() []int {
	url := buildURL(marvel_url, strconv.Itoa(0))
	responseObject := caching(url)
	fmt.Println("responseObject ", responseObject)

	output := make([]int, 0)
	output = concatOutput(responseObject, output)
	fmt.Println("output len ", len(output), " output ", output)

	i := 0
	for {
		i = i + 1
		url := buildURL(marvel_url, strconv.Itoa(i*100))
		responseObject := caching(url)
		output = concatOutput(responseObject, output)
		if len(output) == responseObject.Data.Total {
			break
		}
	}
	return output
}

func caching(url string) Response {
	var responseObject Response
	if exists(url) {
		fmt.Println("cache hit")
		responseObject = get(url)
		return responseObject
	} else {
		fmt.Println("cache miss")
		responseObject, done := queryMarvelApi(url)
		if done {
			//return
		}
		set(url, responseObject)
		return responseObject
	}
}

var PREFIX = "MARVEL_"
func exists(key string) bool {
	var ctx = context.Background()
	return rdb.Exists(ctx, PREFIX+key).Val() != 0
}

func deleteAll() {
	var ctx = context.Background()
	rdb.FlushAll(ctx)
}

func get(key string) Response {
	var ctx = context.Background()

	val, err := rdb.Get(ctx, PREFIX+key).Result()
	if err != nil {
		fmt.Println(err)
	}
	var res Response

	json.Unmarshal([]byte(val), &res)
	return res
}

func set(key string, responseObject Response) interface{} {
	var ctx = context.Background()

	var err error

	b, err := json.Marshal(responseObject)
	err = rdb.Set(ctx, PREFIX+key, b, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func concatOutput(responseObject Response, output []int) []int {
	for _, result := range responseObject.Data.Results {
		output = append(output, result.Id)
	}
	return output
}

func queryMarvelApi(url string) (Response, bool) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return Response{}, true
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	return responseObject, false
}

// GetCharacterById godoc
// @Summary Gets the details of a particular Marvel character
// @Description Serve an endpoint /characters/{characterId} that returns only the id, name and description of the character.
// @Tags Characters
// @Accept  json
// @Produce  json
// @Param characterId path int true "ID of the character"
// @Success 200 {object} Character
// @Failure 400 {object} HTTPError400
// @Failure 404 {object} HTTPError404
// @Failure 500 {object} HTTPError500
// @Router /characters/{characterId} [get]
func getCharacterById (w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	responseObject, _ := queryMarvelApi(buildURL(marvel_url + "/"+vars["id"], strconv.Itoa(0)))
	if responseObject.Code == http.StatusNotFound {
		errorHandler(w, r, responseObject.Code)
		return
	}
	json.NewEncoder(w).Encode(responseObject.Data.Results[0])
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		var response HTTPError404
		response.Code = status
		response.Message = "Not Found"
		fmt.Fprint(w, response)
	}
}