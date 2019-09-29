package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-github/github"
)

func GetKeyBoardList(client *http.Client) []string {
	url := "http://compile.qmk.fm/v1/keyboards"
	var rawJSON json.RawMessage
	var keyboardList []string

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
	}

	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("The HTTP request failed with error %s\n", err)
	}
	rawJSON, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("The JSON could not be parsed with error: %s", err)
	}
	err = json.Unmarshal(rawJSON, &keyboardList)
	if err != nil {
		log.Fatalf("The JSON could not be parsed with error: %s", err)
	}

	return keyboardList
}

func GetKeyMapList(client *github.Client, kbPath string) (keyMapList []string) {
	ctx := context.Background()
	owner := "qmk"
	repo := "qmk_firmware"

	var keyMapPath string

	escapedString := (&url.URL{Path: kbPath}).String()
	keyMapPath = fmt.Sprintf("keyboards/%s/keymaps", escapedString)

	log.Printf("before: %s", kbPath)

	_, directoryContents, _, err := client.Repositories.GetContents(ctx, owner, repo, keyMapPath, nil)
	if err == nil {
		for _, entry := range directoryContents {
			if entry.GetType() == "dir" {
				keyMapList = append(keyMapList, entry.GetName())
			}
		}
	} else {
		log.Print("we have here an error")
		escapedString =  (&url.URL{Path: (path.Dir(escapedString))}).String()
		keyMapPath = fmt.Sprintf("keyboards/%s/keymaps", escapedString)
		_, directoryContents, _, _ := client.Repositories.GetContents(ctx, owner, repo, keyMapPath, nil)
		for _, entry := range directoryContents {
			log.Print(entry)
			if entry.GetType() == "dir" {
				keyMapList = append(keyMapList, entry.GetName())
			}
		}
	}
	return keyMapList
}
