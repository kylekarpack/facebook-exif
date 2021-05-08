package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

var exifMap = make(map[string]Photo)

func main() {
	fmt.Println("Starting...")
	getFiles()
	getPhotos()
}

func readFile(filename string) Album {
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
		var parts = strings.Split(photo.URI, "/")
		var name = parts[len(parts)-1]
		exifMap[name] = photo
	}
}

func processPhoto(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	exifData, err := exif.Decode(file)

	fmt.Println(exifData)
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
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}

func getPhotos() {

	photos, err := filepath.Glob("./photos_and_videos/**/*.jpg")

	if err != nil {
		log.Println(err)
		return
	}

	for _, photo := range photos {

	}

}
