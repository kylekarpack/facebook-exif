package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
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
	// get last modified time
	file, err := os.Stat(path)

	if err != nil {
		fmt.Println(err)
	}

	modifiedtime := file.ModTime()

	fmt.Println("Last modified time : ", modifiedtime)

	// get current timestamp

	currenttime := time.Now().Local()

	fmt.Println("Current time : ", currenttime.Format("2006-01-02 15:04:05 +0800"))

	// change both atime and mtime to currenttime

	err = os.Chtimes(path, currenttime, currenttime)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Changed the file time information")

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
