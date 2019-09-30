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
	"time"

	"github.com/google/go-github/github"
)

var httpClient = &http.Client{
	Timeout: time.Second * 2,
}

var githubClient = github.NewClient(httpClient)

func GetKeyBoardList() []string {
	url := "http://compile.qmk.fm/v1/keyboards"
	var rawJSON json.RawMessage
	var keyboardList []string

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
	}

	res, err := httpClient.Do(req)
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

func GetKeyMapList(kbPath string) (keyMapList []string) {
	ctx := context.Background()
	owner := "qmk"
	repo := "qmk_firmware"

	var keyMapPath string

	escapedString := (&url.URL{Path: kbPath}).String()
	keyMapPath = fmt.Sprintf("keyboards/%s/keymaps", escapedString)

	log.Printf("before: %s", kbPath)

	_, directoryContents, _, err := githubClient.Repositories.GetContents(ctx, owner, repo, keyMapPath, nil)
	if err == nil {
		keyMapList = _getKeymaps(directoryContents)
	} else {
		// for the outlier case where the keymap is kept in the parent directory
		escapedString = (&url.URL{Path: (path.Dir(escapedString))}).String()
		keyMapPath = fmt.Sprintf("keyboards/%s/keymaps", escapedString)
		_, directoryContents, _, secondErr := githubClient.Repositories.GetContents(ctx, owner, repo, keyMapPath, nil)
		if secondErr == nil {
			keyMapList = _getKeymaps(directoryContents)
		} else {
			for _, subEntry := range directoryContents {
				keyMapPath = fmt.Sprintf("keyboards/%s/%s/keymaps", escapedString, subEntry.GetName())
				_, directoryContents, _, _ = githubClient.Repositories.GetContents(ctx, owner, repo, keyMapPath, nil)
				keyMapList = _getKeymaps(directoryContents)
			}
		}
	}
	return keyMapList
}
// Sample code for streaming output from console to qtextedit
/*
    stream := run(input.Text())
            go func() {
                    for line := range stream {
                            output.Append(line)
                    }
            }()
    })

    // ...
}

z*/


func _getKeymaps(directoryContents []*github.RepositoryContent) (keyMapList []string) {
	for _, entry := range directoryContents {
		if entry.GetType() == "dir" {
			keyMapList = append(keyMapList, entry.GetName())
		}
	}
	return keyMapList
}
