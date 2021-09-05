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

func Init (path string) {
	ReadFile(&Cfg, path)
	fmt.Printf("\nLoaded configuration %+v\n\n", Cfg)
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func GetCharacterIdUrl(id string) string {
	return constant.MARVEL_URL + "/" + id
}

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


func processError(err error) {
	fmt.Println(err)
	os.Exit(2)
}

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

func GetExpirationBasedOnStrategy() time.Duration {
	if *Strategy == "TTL"{
		return Cfg.Marvel.TTL * time.Minute
	}

	return 0
}