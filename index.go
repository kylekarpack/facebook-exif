package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var m map[string]string

func main() {
	fmt.Println("Starting...")
	getFiles()
}

func readFile(filename string) Album {
	fmt.Println(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var album Album
	json.Unmarshal([]byte(content), &album)

	return album
}

func processAlbum(album Album) {
	for _, photo := range album.Photos {
		//photo.URI
	}
}

func getFiles() {
	err := filepath.Walk("./photos_and_videos/album",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if strings.HasSuffix(path, ".json") {
				album := readFile(path)
				processAlbum(album)

				fmt.Println(album)
			}

			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
