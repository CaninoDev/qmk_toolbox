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

func GetKeyBoardList(client http.Client) []string {
	url := "http://compile.qmk.fm/v1/keyboards"
	var rawJSON json.RawMessage
	var keyboardList []string

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Print(err)
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	}
	rawJSON, err = ioutil.ReadAll(res.Body)
	if err != nil {
		log.Printf("The JSON could not be parsed with error: %s", err)
	}
	err = json.Unmarshal(rawJSON, &keyboardList)
	if err != nil {
		log.Printf("The JSON could not be parsed with error: %s", err)
	}

	return keyboardList
}

func GetKeyMapList(ctx context.Context, client *github.Client, kbPath string) (keyMapList []string, err error) {
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
		log.Print(keyMapPath)
		_, directoryContents, _, _ := client.Repositories.GetContents(ctx, owner, repo, keyMapPath, nil)
		for _, entry := range directoryContents {
			log.Print(entry)
			if entry.GetType() == "dir" {
				keyMapList = append(keyMapList, entry.GetName())
			}
		}
	}
	return keyMapList, nil
}

// func GetList(ctx context.Context, client github.Client, owner string, repo string, pth string) (list []string) {
// 	_, directoryContents, _, _ := client.Repositories.GetContents(ctx, owner, repo, pth, nil)
// 	for _, content := range directoryContents {
// 		_, subDirectoryContents, _, _ := client.Repositories.GetContents(ctx, owner, repo, subPath, nil)
// 		for i := range subDirectoryContents {
// 			if *subDirectoryContents[i].Name == "keymaps" {

// 			}
// 		}
// 	}

// }
//var myClient = &http.Client{Timeout: 10 * time.Second}
//
//type RepositoryContent struct {
//	Type *string `json:"type,omitempty"`
//	// Target is only set if the type is "symlink" and the target is not a normal file.
//	// If Target is set, Path will be the symlink path.
//	Target   *string `json:"target,omitempty"`
//	Encoding *string `json:"encoding,omitempty"`
//	Size     *int    `json:"size,omitempty"`
//	Name     *string `json:"name,omitempty"`
//	Path     *string `json:"path,omitempty"`
//	// Content contains the actual file content, which may be encoded.
//	// Callers should call GetContent which will decode the content if
//	// necessary.
//	Content     *string `json:"content,omitempty"`
//	SHA         *string `json:"sha,omitempty"`
//	URL         *string `json:"url,omitempty"`
//	GitURL      *string `json:"git_url,omitempty"`
//	HTMLURL     *string `json:"html_url,omitempty"`
//	DownloadURL *string `json:"download_url,omitempty"`
//}
//
//func GetList(ctx context.Context, ) ([]string, error) {
//	ctx = context.Background()
//	url := "https://api.github.com/repos/qmk/qmk_firmware/contents/keyboards"
//	var rawJSON json.RawMessage
//	var directoryContent []RepositoryContent
//
//	res, err := myClient.Get(url)
//	if err != nil {
//		return (nil, err)
//	}
//	defer res.Body.Close()
//
//	rawJSON, err = []byte(ioutil.ReadAll(res.Body))
//	if (err != nil) {
//		return (nil, err)
//	}
//
//	directoryUnmarshalError := json.Unmarshal(res.Body, &directoryContent)
//	if directoryUnmarshalError == nil {
//
//	}
//}
