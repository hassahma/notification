// Service methods used by controller to fetch and process data from external api.
package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"github.com/marvel/utils"
	"github.com/marvel/model"
	"github.com/marvel/cache"
	"github.com/marvel/constant"
)

// Invalidates the cache by calling flushall
func InvalidateAndRefresh() {
	fmt.Println("Invalidating and Prefetching cache")
	cache.DeleteAll()
	fetch()
}

// calls the utility method to fetch data from external api and then process the response to build the list of character IDs.
func fetch() []int {
	url := utils.BuildURL(constant.MARVEL_URL, strconv.Itoa(0))
	responseObject := caching(url)
	output := make([]int, 0)
	output = concatOutput(responseObject, output)

	if utils.Loop != 1 {
		return output
	}

	i := 0
	for {
		i = i + 1
		url := utils.BuildURL(constant.MARVEL_URL, strconv.Itoa(i*100))
		responseObject := caching(url)
		output = concatOutput(responseObject, output)
		if len(output) == responseObject.Data.Total {
			break
		}
	}
	return output
}

// Checks the cache if the key exists and return the response from the cache otherwise fetches the data from external api and store it in cache.
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
			fmt.Println("Failed to access downstream service ", responseObject)
		}
		cache.Set(url, responseObject)
		return responseObject
	}
}

// utility method to append the IDs from Response object to the integer array.
func concatOutput(responseObject model.Response, output []int) []int {
	for _, result := range responseObject.Data.Results {
		output = append(output, result.Id)
	}
	return output
}

// Makes a http GET request to external api based on the url passed.
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
