package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
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
			t := time.Unix(int64(val.CreationTimestamp), 0)
			strDate := t.Format(time.RFC3339)
			cmd := exec.Command("exiftool", "-overwrite_original", "-AllDates="+strDate, filepath)
			cmd.Run()
			fmt.Println("Fixed date for", filename, "to", strDate)
		} else {
			fmt.Println("Could not find ", filename)
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

func getFiles() map[string]Photo {

	albums, err := filepath.Glob("./photos_and_videos/album/*.json")

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

func getPhotos() []string {

	photos, err := filepath.Glob("./photos_and_videos/**/*.jpg")
	if err != nil {
		log.Panic(err)
	}

	return photos

}
