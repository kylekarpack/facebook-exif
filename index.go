package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var exifMap = make(map[string]Photo)

func main() {
	fmt.Println("Starting...")
	run()
	fmt.Println("Complete")
}

func checkExiftool() {
	// Check for exiftool
	cmd := exec.Command("exiftool", "-ver")
	err := cmd.Run()
	if err != nil {
		log.Fatal("exiftool is required to run this program. Please install with \"sudo apt-get install exiftool\"")
	}
}

func run() {
	exifMap := getFiles()
	fmt.Print("This many", len(exifMap))
	photos := getPhotos()
	fixDates(photos, exifMap)
}

func fixDates(photos []string, exifMap map[string]Photo) {
	for _, filepath := range photos {
		filename := getFilenameFromPath(filepath)
		if val, ok := exifMap[filename]; ok {
			i, err := strconv.ParseInt((string)(val.CreationTimestamp), 10, 64)
			if err != nil {
				log.Panic(err)
			}
			tm := time.Unix(i, 0)
			fmt.Println(tm)
			// cmd := exec.Command("exiftool", "-AllDates=2021:05:02 21:36:17-04:00", filepath)
			// cmd.Run()
			fmt.Println("Fixed metadata for", filepath, "to", tm)
		} else {
			log.Fatal("Could not find ", filename)
		}

	}
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

func processAlbum(album Album, exifMap map[string]Photo) {
	for _, photo := range album.Photos {
		name := getFilenameFromPath(photo.URI)
		exifMap[name] = photo
	}
}

func getFilenameFromPath(path string) string {
	var parts = strings.Split(path, "/")
	var name = parts[len(parts)-1]
	return name
}

func getFiles() map[string]Photo {

	albums, err := filepath.Glob("./photos_and_videos/albums/**/*.json")

	if err != nil {
		log.Panic(err)
	}

	var exifMap = make(map[string]Photo)

	for _, path := range albums {
		album := readFile(path)
		processAlbum(album, exifMap)
	}

	return exifMap
}

func getPhotos() []string {

	photos, err := filepath.Glob("./photos_and_videos/**/*.jpg")
	if err != nil {
		log.Panic(err)
	}

	return photos

}
