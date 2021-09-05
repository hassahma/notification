package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/marvel/errorhandler"
	"github.com/marvel/utils"
	"github.com/marvel/model"
)

var marvel_url = "https://gateway.marvel.com/v1/public/characters"

var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

var characters []model.Character

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
func GetAllCharacters(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all characters")
	json.NewEncoder(w).Encode(fetch())
}

func InvalidateAndRefresh() {
	fmt.Println("Invalidating and Prefetching cache")
	deleteAll()
	fetch()
}

func fetch() []int {
	url := utils.BuildURL(marvel_url, strconv.Itoa(0))
	responseObject := caching(url)
	output := make([]int, 0)
	output = concatOutput(responseObject, output)

	i := 0
	for {
		i = i + 1
		url := utils.BuildURL(marvel_url, strconv.Itoa(i*100))
		responseObject := caching(url)
		output = concatOutput(responseObject, output)
		if len(output) == responseObject.Data.Total {
			break
		}
	}
	return output
}

func caching(url string) model.Response {
	var responseObject model.Response
	if exists(url) {
		fmt.Println("Cache hit")
		responseObject = get(url)
		return responseObject
	} else {
		fmt.Println("Cache miss")
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

func get(key string) model.Response {
	var ctx = context.Background()

	val, err := rdb.Get(ctx, PREFIX+key).Result()
	if err != nil {
		fmt.Println(err)
	}
	var res model.Response

	json.Unmarshal([]byte(val), &res)
	return res
}

func set(key string, responseObject model.Response) interface{} {
	var ctx = context.Background()

	var err error

	b, err := json.Marshal(responseObject)
	err = rdb.Set(ctx, PREFIX+key, b, 0).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func concatOutput(responseObject model.Response, output []int) []int {
	for _, result := range responseObject.Data.Results {
		output = append(output, result.Id)
	}
	return output
}

func queryMarvelApi(url string) (model.Response, bool) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return model.Response{}, true
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject model.Response
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
func GetCharacterById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	responseObject, _ := queryMarvelApi(utils.BuildURL(marvel_url+"/"+vars["id"], strconv.Itoa(0)))
	if responseObject.Code == http.StatusNotFound {
		errorhandler.ErrorHandler(w, r, responseObject.Code)
		return
	}
	json.NewEncoder(w).Encode(responseObject.Data.Results[0])
}
