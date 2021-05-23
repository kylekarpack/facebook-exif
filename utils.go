package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"strings"
)

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
