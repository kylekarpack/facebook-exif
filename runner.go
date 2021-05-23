package main

import (
	"fmt"
	"time"
)

func run(dir string, dryRun bool) {
	start := time.Now()
	fmt.Println("Starting...")
	exifMap := getFiles(dir)
	photos := getPhotos(dir)
	if dryRun {
		logInfo(photos, exifMap)
	} else {
		fixDates(photos, exifMap)
	}
	fmt.Printf("Complete in %v\n", time.Since(start))
}

// Dry run: log info without modifying files
func logInfo(photos []string, exifMap map[string]Photo) {
	fmt.Printf("%v photos found\n", len(photos))
}

// Full run: modify photo metadata
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
