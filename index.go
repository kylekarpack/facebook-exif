package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
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
		log.Panic(err)
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
	cmd := exec.Command("exiftool", "-AllDates=2021:05:02 21:36:17-04:00", path)
	cmd.Run()
	fmt.Println("Fixed metadata for", path)

}

func getFiles() {

	albums, err := filepath.Glob("./photos_and_videos/albums/**/*.json")

	if err != nil {
		log.Panic(err)
	}

	for _, path := range albums {
		album := readFile(path)
		processAlbum(album)
	}
}

func getPhotos() {

	photos, err := filepath.Glob("./photos_and_videos/**/*.jpg")
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(photos[0])
	processPhoto(photos[0])
	// for _, photo := range photos {
	// 	processPhoto(photo)
	// }

}
