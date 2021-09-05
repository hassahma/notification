package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/marvel/utils"
	"github.com/marvel/cache"
	"io"
	"strconv"
	"github.com/gorilla/mux"
)

func setup() {
	utils.Loop = 0
	utils.Init("../config/config.yml")
	cache.Init()
}

func TestGetAllCharacters(t *testing.T) {

	setup()

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/characters", nil)

	GetAllCharacters(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if wr.Body.String() == "" {
		t.Errorf(
			`response body "%s" is empty!"`,
			wr.Body.String(),
		)
	}
}

func TestGetCharactersById(t *testing.T) {

	setup()

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/characters", nil)

	GetAllCharacters(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if wr.Body.String() == "" {
		t.Errorf(
			`response body "%s" is empty!"`,
			wr.Body.String(),
		)
	}

	resp := wr.Result()
	body, _ := io.ReadAll(resp.Body)
	res := make([]int, 0)
	json.Unmarshal([]byte(body), &res)

	id := strconv.Itoa(res[0])

	req = httptest.NewRequest(http.MethodGet, "/characters/", nil)

	vars := map[string]string{
		"id": id,
	}
	req = mux.SetURLVars(req, vars)
	GetCharacterById(wr, req)
	if wr.Code != http.StatusOK {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}

	if wr.Body.String() == "" {
		t.Errorf(
			`response body "%s" is empty!"`,
			wr.Body.String(),
		)
	}
}

func TestGetCharactersById404(t *testing.T) {

	setup()

	wr := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/characters/", nil)

	vars := map[string]string{
		"id": "non_existing_id",
	}
	req = mux.SetURLVars(req, vars)
	GetCharacterById(wr, req)
	if wr.Code != http.StatusNotFound {
		t.Errorf("got HTTP status code %d, expected 200", wr.Code)
	}
}
