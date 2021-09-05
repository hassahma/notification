// utils package that provides various utility methods used within the application
package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"fmt"
	"log"
	"github.com/marvel/constant"
	"github.com/marvel/model"
	"os"
	"time"
	"gopkg.in/yaml.v2"
)

var Cfg model.Config
var defaultstrategy = "PREFETCH"
var Strategy *string = &defaultstrategy
var Loop = 1

// initializes the application
func Init (path string) {
	ReadFile(&Cfg, path)
	fmt.Printf("\nLoaded configuration %+v\n\n", Cfg)
}

// generates the MD5 hash of a string
func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

// Generates the url string for /characters/{id} api call
func GetCharacterIdUrl(id string) string {
	return constant.MARVEL_URL + "/" + id
}

// builds the url string from base url and offset
func BuildURL(urlStr string, offset string) string {
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}

	url.Scheme = constant.SCHEME
	queryparams := url.Query()
	queryparams.Set("ts", Cfg.Marvel.Timestamp)
	queryparams.Set("apikey", Cfg.Marvel.Apikey)
	queryparams.Set("hash", GetMD5Hash(Cfg.Marvel.Timestamp + Cfg.Marvel.Privatekey + Cfg.Marvel.Apikey))
	queryparams.Set("limit", Cfg.Marvel.Limit)
	queryparams.Set("offset", offset)

	url.RawQuery = queryparams.Encode()
	fmt.Println("Query is " + url.String())
	return url.String()
}

// handles the error from reading config file
func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

// reads the configuration file
func ReadFile(cfg *model.Config, path string) {
	f, err := os.Open(path)
	if err != nil {
		processError(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(cfg)
	if err != nil {
		processError(err)
	}
}

// Returns the cache TTL based on the chosen strategy. For PREFETCH '0' is returned which means the key has no expiry.
func GetExpirationBasedOnStrategy() time.Duration {
	if *Strategy == "TTL"{
		return Cfg.Marvel.TTL * time.Minute
	}

	return 0
}