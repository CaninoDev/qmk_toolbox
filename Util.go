package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/graphql-go/graphql"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	res, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}
