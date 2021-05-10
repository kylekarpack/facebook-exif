package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	jpeg "github.com/dsoprea/go-jpeg-image-structure/v2"
)

var exifMap = make(map[string]Photo)

func main() {
	fmt.Println("Starting...")
	exifMap := getFiles()
	photos := getPhotos()
	fixDates(photos, exifMap)
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

func fixDates(photos []string, exifMap map[string]Photo) {
	for _, filepath := range photos[0:10] {
		filename := getFilenameFromPath(filepath)
		if val, ok := exifMap[filename]; ok {
			t := time.Unix(int64(val.CreationTimestamp), 0)
			setPhotoDate(filepath, t)
			fmt.Println("Fixed date for", filename, "to", t.Format(time.RFC3339))
		} else {
			fmt.Println("Could not find ", filename)
		}
	}
}

func setExifTag(rootIB *exif.IfdBuilder, ifdPath, tagName, tagValue string) error {
	ifdIb, err := exif.GetOrCreateIbFromRootIb(rootIB, ifdPath)
	if err != nil {
		return fmt.Errorf("Failed to get or create IB: %v", err)
	}

	if err := ifdIb.SetStandardWithName(tagName, tagValue); err != nil {
		return fmt.Errorf("failed to set DateTime tag: %v", err)
	}

	return nil
}

func setPhotoDate(filepath string, t time.Time) error {
	parser := jpeg.NewJpegMediaParser()
	intfc, err := parser.ParseFile(filepath)
	if err != nil {
		return fmt.Errorf("Failed to parse JPEG file: %v", err)
	}

	sl := intfc.(*jpeg.SegmentList)

	rootIb, err := sl.ConstructExifBuilder()
	if err != nil {
		fmt.Println("No EXIF; creating it from scratch")

		im, err := exifcommon.NewIfdMappingWithStandard()
		if err != nil {
			return fmt.Errorf("Failed to create new IFD mapping with standard tags: %v", err)
		}
		ti := exif.NewTagIndex()
		if err := exif.LoadStandardTags(ti); err != nil {
			return fmt.Errorf("Failed to load standard tags: %v", err)
		}

		rootIb = exif.NewIfdBuilder(im, ti, exifcommon.IfdStandardIfdIdentity,
			exifcommon.EncodeDefaultByteOrder)
		rootIb.AddStandardWithName("ProcessingSoftware", "photos-uploader")
	}

	// Form our timestamp string
	ts := exifcommon.ExifFullTimestampString(t)

	// Set DateTime
	ifdPath := "IFD0"
	if err := setExifTag(rootIb, ifdPath, "DateTime", ts); err != nil {
		return fmt.Errorf("Failed to set tag %v: %v", "DateTime", err)
	}

	// Set DateTimeOriginal
	ifdPath = "IFD/Exif"
	if err := setExifTag(rootIb, ifdPath, "DateTimeOriginal", ts); err != nil {
		return fmt.Errorf("Failed to set tag %v: %v", "DateTimeOriginal", err)
	}

	// Update the exif segment.
	if err := sl.SetExif(rootIb); err != nil {
		return fmt.Errorf("Failed to set EXIF to jpeg: %v", err)
	}

	// Write the modified file
	b := new(bytes.Buffer)
	if err := sl.Write(b); err != nil {
		return fmt.Errorf("Failed to create JPEG data: %v", err)
	}

	// Save the file
	if err := ioutil.WriteFile(filepath, b.Bytes(), 0644); err != nil {
		return fmt.Errorf("Failed to write JPEG file: %v", err)
	}

	return nil
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
