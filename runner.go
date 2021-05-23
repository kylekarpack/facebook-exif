package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func run(dir string, dryRun bool) {
	fmt.Println("Starting...")
	exifMap := getFiles(dir)
	photos := getPhotos(dir)
	if dryRun {

	} else {
		fixDates(photos, exifMap)
	}
	fmt.Println("Complete")
}

func fixDates(photos []string, exifMap map[string]Photo) {
	for i, filepath := range photos {
		filename := getFilenameFromPath(filepath)
		if val, ok := exifMap[filename]; ok {
			t := time.Unix(int64(val.CreationTimestamp), 0)
			setPhotoDate(filepath, t)
			fmt.Printf("(%v of %v) Fixed date for %v to %v\n", i, len(photos), filename, t.Format(time.RFC3339))
		} else {
			fmt.Printf("Could not find file: %v\n", filename)
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

func getFilenameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

func getFiles(dir string) map[string]Photo {

	albums, err := filepath.Glob(path.Join(dir, "/album/*.json"))

	if err != nil {
		log.Panic(err)
	}

	exifMap := make(map[string]Photo)

	for _, path := range albums {
		album := readFile(path)
		for _, photo := range album.Photos {
			name := getFilenameFromPath(photo.URI)
			exifMap[name] = photo
		}
	}
	return exifMap
}

func getPhotos(dir string) []string {

	photos, err := filepath.Glob(path.Join(dir, "/**/*.jpg"))
	if err != nil {
		log.Panic(err)
	}

	return photos
}
