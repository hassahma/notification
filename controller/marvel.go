package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/marvel/errorhandler"
	"github.com/marvel/utils"
)

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
	responseObject, _ := queryMarvelApi(utils.BuildURL(utils.GetCharacterIdUrl(vars["id"]), strconv.Itoa(0)))
	if responseObject.Code == http.StatusNotFound {
		errorhandler.ErrorHandler(w, r, responseObject.Code)
		return
	}
	json.NewEncoder(w).Encode(responseObject.Data.Results[0])
}
