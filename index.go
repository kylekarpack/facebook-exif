package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	exif "github.com/dsoprea/go-exif/v3"
	jpegstructure "github.com/dsoprea/go-jpeg-image-structure/v2"
	log "github.com/dsoprea/go-logging"
)

var exifMap = make(map[string]Photo)

func main() {
	fmt.Println("Starting...")
	getFiles()
	getPhotos()
}

func readFile(filename string) Album {
	content, err := ioutil.ReadFile(filename)
	log.PanicIf(err)

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
	jmp := jpegstructure.NewJpegMediaParser()

	intfc, err := jmp.ParseFile(path)
	log.PanicIf(err)

	sl := intfc.(*jpegstructure.SegmentList)

	// Update the CameraOwnerName tag.

	rootIb, err := sl.ConstructExifBuilder()
	log.PanicIf(err)

	ifdPath := "IFD/Exif"

	ifdIb, err := exif.GetOrCreateIbFromRootIb(rootIb, ifdPath)
	log.PanicIf(err)

	err = ifdIb.SetStandardWithName("CameraOwnerName", "TestOwner")
	log.PanicIf(err)

	// Update the exif segment.

	err = sl.SetExif(rootIb)
	log.PanicIf(err)

	b := new(bytes.Buffer)

	err = sl.Write(b)
	log.PanicIf(err)

	// Validate.

	d := b.Bytes()

	intfc, err = jmp.ParseBytes(d)
	log.PanicIf(err)

	sl = intfc.(*jpegstructure.SegmentList)

	_, _, exifTags, err := sl.DumpExif()
	log.PanicIf(err)

	for _, et := range exifTags {
		if et.IfdPath == "IFD/Exif" && et.TagName == "CameraOwnerName" {
			fmt.Printf("Value: [%s]\n", et.FormattedFirst)
			break
		}
	}

}

func getFiles() {

	albums, err := filepath.Glob("./photos_and_videos/albums/**/*.json")

	log.PanicIf(err)

	for _, path := range albums {
		album := readFile(path)
		processAlbum(album)
	}
}

func getPhotos() {

	photos, err := filepath.Glob("./photos_and_videos/**/*.jpg")
	log.PanicIf(err)

	fmt.Println(photos[0])
	processPhoto(photos[0])
	// for _, photo := range photos {
	// 	processPhoto(photo)
	// }

}
