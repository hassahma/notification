// tests for utility functions.
package utils

import (
	"time"
	"testing"
	"github.com/marvel/constant"
	"github.com/marvel/testutils"
)

// tests the generation of md5 hash.
func TestGetMD5Hash(t *testing.T) {
	testutil.AssertEqual(t, GetMD5Hash("dummy"), "275876e34cf609db118f3d84b799a790")
}

// tests the generation of url
func TestGetCharacterIdUrl(t *testing.T) {
	testutil.AssertEqual(t, GetCharacterIdUrl("1234"), constant.MARVEL_URL+"/1234")
}

// tests the TTL time based on chosen strategy.
func TestGetExpirationBasedOnStrategy(t *testing.T) {
	Init("../config/config.yml")
	testutil.AssertEqual(t, GetExpirationBasedOnStrategy(), time.Duration(0)*time.Minute)
	var defaultstrategy = "TTL"
	Strategy = &defaultstrategy
	testutil.AssertEqual(t, GetExpirationBasedOnStrategy(), time.Duration(2)*time.Minute)
}

// tests build url string
func TestBuildURL(t *testing.T) {
	expecting := constant.MARVEL_URL + "?apikey=" + Cfg.Marvel.Apikey + "&hash=" + GetMD5Hash(Cfg.Marvel.Timestamp+Cfg.Marvel.Privatekey+Cfg.Marvel.Apikey) + "&limit=" + Cfg.Marvel.Limit + "&offset=100&ts=" + Cfg.Marvel.Timestamp
	testutil.AssertEqual(t, BuildURL(constant.MARVEL_URL, "100"), expecting)
}

// tests reading config file
func TestReadFile(t *testing.T) {
	ReadFile(&Cfg, "../config/config.yml")
	testutil.AssertEqual(t, Cfg.Marvel.Timestamp, "1")
	testutil.AssertEqual(t, Cfg.Marvel.Limit, "100")
}
