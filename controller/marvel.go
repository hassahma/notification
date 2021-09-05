package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/marvel/errorhandler"
	"github.com/marvel/utils"
	"github.com/marvel/model"
	"github.com/marvel/cache"
)

var marvel_url = "https://gateway.marvel.com/v1/public/characters"

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
	fmt.Println("Get character by ID")
	vars := mux.Vars(r)
	responseObject, _ := queryMarvelApi(utils.BuildURL(marvel_url+"/"+vars["id"], strconv.Itoa(0)))
	if responseObject.Code == http.StatusNotFound {
		errorhandler.ErrorHandler(w, r, responseObject.Code)
		return
	}
	json.NewEncoder(w).Encode(responseObject.Data.Results[0])
}

func InvalidateAndRefresh() {
	fmt.Println("Invalidating and Prefetching cache")
	cache.DeleteAll()
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
	if cache.Exists(url) {
		fmt.Println("Cache hit")
		responseObject = cache.Get(url)
		return responseObject
	} else {
		fmt.Println("Cache miss")
		responseObject, done := queryMarvelApi(url)
		if done {
			//return
		}
		cache.Set(url, responseObject)
		return responseObject
	}
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
