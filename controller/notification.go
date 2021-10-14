// package implements the controller methods for the REST api.
package controller

import (
	"fmt"
	"net/http"
)

// PostNotify godoc
// @Summary Test the notification
// @Description todo
// @Tags Notification
// @Accept  json
// @Produce  json
// @Success 201
// @Failure 400 {object} model.HTTPError400
// @Failure 500 {object} model.HTTPError500
// @Param url query string true "The notification url"
// @Param customerId query int true "The id of the customer"
// @Param body_param body string true "The payload"
// @Router /notification/subscribe [post]
func PostNotify(w http.ResponseWriter, r *http.Request) {
	saveCustomer(w, r)
	w.WriteHeader(http.StatusCreated)
}

// PostNotifyTest godoc
// @Summary Test the notification
// @Description todo
// @Tags Notification
// @Accept  json
// @Produce  json
// @Success 201
// @Failure 400 {object} model.HTTPError400
// @Failure 500 {object} model.HTTPError500
// @Router /notification/test [post]
func PostNotifyTest(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostNotifyTest called")
	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(fetch())
}

// PostNotifyActivate godoc
// @Summary Activate the notification
// @Description todo
// @Tags Notification
// @Accept  json
// @Produce  json
// @Success 201
// @Failure 400 {object} model.HTTPError400
// @Failure 500 {object} model.HTTPError500
// @Router /notification/activate [post]
func PostNotifyActivate(w http.ResponseWriter, r *http.Request) {
	fmt.Println("PostNotifyActivate called")
	w.WriteHeader(http.StatusCreated)
	//json.NewEncoder(w).Encode(fetch())
}

/*
// GetCharacterById godoc
// @Summary Gets the details of a particular Marvel character
// @Description Serve an endpoint /characters/{characterId} that returns only the id, name and description of the character.
// @Tags Characters
// @Accept  json
// @Produce  json
// @Param characterId path int true "ID of the character"
// @Success 200 {object} model.Character
// @Failure 400 {object} model.HTTPError400
// @Failure 404 {object} model.HTTPError404
// @Failure 500 {object} model.HTTPError500
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
*/
