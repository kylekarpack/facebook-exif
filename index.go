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
	"github.com/rwcarlsen/goexif/tiff"
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

type Printer struct{}

func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%40s: %s\n", name, tag)
	return nil
}

func processPhoto(path string) {
	fmt.Println(path)
	file, err := os.Open("./" + path)
	if err != nil {
		fmt.Println("Badfasdfsadf")
		log.Fatal(err)
	}

	exifData, err := exif.Decode(file)

	if err != nil {
		fmt.Println("Badfasdfsadf 123123")

		log.Fatal(err)
	}

	var p Printer

	exifData.Walk(p)

}

func getFiles() {

	albums, err := filepath.Glob("./photos_and_videos/albums/**/*.json")

	if err != nil {
		log.Fatal(err)
	}

	for _, path := range albums {
		album := readFile(path)
		processAlbum(album)
	}
}

func getPhotos() {

	photos, err := filepath.Glob("./photos_and_videos/**/*.jpg")

	if err != nil {
		log.Fatal(err)
	}

	for _, photo := range photos {
		processPhoto(photo)
	}

}
