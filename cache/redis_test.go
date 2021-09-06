package cache

import (
	"testing"
	"github.com/marvel/utils"
	"github.com/marvel/testutils"
	"github.com/marvel/model"
	"reflect"
)

func setup() {
	utils.Loop = 0
	utils.Init("../config/config.yml")
	Init()
	DeleteAll()
}

func teardown() {
	DeleteAll()
}

// tests the generation of md5 hash.
func TestExists(t *testing.T) {
	setup()
	key := "dummy"
	var val = model.Response{}
	testutil.AssertEqual(t, Exists(key), false)
	Set(key, val)
	testutil.AssertEqual(t, Exists(key), true)
	teardown()
}

// tests the generation of md5 hash.
func TestGet(t *testing.T) {
	setup()
	key := "dummy"
	var val = model.Response{}
	Set(key, val)
	if !reflect.DeepEqual(Get(key), val) {
		t.Fatalf("cache value is not the same")
	}
	teardown()
}

// tests the generation of md5 hash.
func TestSet(t *testing.T) {
	setup()
	key := "dummy"
	var val = model.Response{}
	testutil.AssertEqual(t, Exists(key), false)
	Set(key, val)
	if !reflect.DeepEqual(Get(key), val) {
		t.Fatalf("cache value is not the same")
	}
	teardown()
}
