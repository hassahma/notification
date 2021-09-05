package utils

import (
	"crypto/md5"
	"encoding/hex"
	"net/url"
	"fmt"
	"log"
)

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func BuildURL(urlStr string, offset string) string {
	ts := "1"
	apikey := "17c400eff8dafac0184fb02420750089"
	privateKey := "1d724a31e223d1fd04a442b129c66b4b6528b360"
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Fatal(err)
	}

	url.Scheme = "https"
	queryparams := url.Query()
	queryparams.Set("ts", ts)
	queryparams.Set("apikey", apikey)
	queryparams.Set("hash", GetMD5Hash(ts+privateKey+apikey))
	queryparams.Set("limit", "100")
	queryparams.Set("offset", offset)

	url.RawQuery = queryparams.Encode()
	fmt.Println("Query is " + url.String())
	return url.String()
}