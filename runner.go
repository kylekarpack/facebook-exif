package main

import (
	"fmt"
	"time"
)

// Run the CLI application
func run(dir string, dryRun bool) {
	defer duration(track("Completed in"))
	fmt.Println("Starting...")
	metadataMap := getMetadata(dir)
	photos := getPhotos(dir)
	if dryRun {
		logInfo(photos, metadataMap)
	} else {
		fixDates(photos, metadataMap)
	}
}

// Dry run: log info without modifying files
func logInfo(photos []string, metadataMap map[string]Photo) {
	fmt.Printf("%v photos found\n", len(photos))
	fmt.Printf("%v metadata entries found\n", len(metadataMap))
}

// Full run: modify photo metadata
func fixDates(photos []string, metadataMap map[string]Photo) {
	for i, filepath := range photos {
		filename := getFilenameFromPath(filepath)
		if val, ok := metadataMap[filename]; ok {
			t := time.Unix(int64(val.CreationTimestamp), 0)
			setPhotoDate(filepath, t)
			fmt.Printf("(%v of %v) Fixed date for %v to %v\n", i, len(photos), filename, t.Format(time.RFC3339))
		} else {
			fmt.Printf("Could not find file: %v\n", filename)
		}
	}
}
