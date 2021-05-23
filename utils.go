package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"
	"time"
)

// Parse a JSON file and retrieve its metadata
func readFile(filename string) Album {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var album Album
	json.Unmarshal([]byte(content), &album)

	return album
}

// Get the filename (last part) from a full path string
func getFilenameFromPath(path string) string {
	parts := strings.Split(path, "/")
	return parts[len(parts)-1]
}

// Get metadata as a map of file names to metatdata objects
func getMetadata(dir string) map[string]Photo {

	albums, err := filepath.Glob(path.Join(dir, "/album/*.json"))

	if err != nil {
		log.Fatal(err)
	}

	if len(albums) == 0 {
		log.Fatalf("No metadata found in directory %v\n", dir)
	}

	metadataMap := make(map[string]Photo)

	for _, path := range albums {
		album := readFile(path)
		for _, photo := range album.Photos {
			name := getFilenameFromPath(photo.URI)
			metadataMap[name] = photo
		}
	}
	return metadataMap
}

// Get all photos in the directory recursively
func getPhotos(dir string) []string {

	photos, err := filepath.Glob(path.Join(dir, "/**/*.jpg"))
	if err != nil {
		log.Fatal(err)
	}

	if len(photos) == 0 {
		log.Fatalf("No photos found in directory %v\n", dir)
	}

	return photos
}

// Track the execution time of a function
func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

// Print the execution time of a function
func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))
}
