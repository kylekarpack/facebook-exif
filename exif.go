package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	exif "github.com/dsoprea/go-exif/v3"
	exifcommon "github.com/dsoprea/go-exif/v3/common"
	jpeg "github.com/dsoprea/go-jpeg-image-structure/v2"
	riimage "github.com/dsoprea/go-utility/v2/image"
)

// Set a given tag value on an exif object
func setExifTag(rootIB *exif.IfdBuilder, ifdPath, tagName, tagValue string) {
	ifdIb, err := exif.GetOrCreateIbFromRootIb(rootIB, ifdPath)
	if err != nil {
		log.Fatalf("Failed to get or create IB: %v", err)
	}

	if err := ifdIb.SetStandardWithName(tagName, tagValue); err != nil {
		log.Fatalf("failed to set %v tag to %v: %v", tagName, tagValue, err)
	}
}

func getMediaContext(filepath string) riimage.MediaContext {
	parser := jpeg.NewJpegMediaParser()
	context, err := parser.ParseFile(filepath)
	if err != nil {
		log.Fatalf("Failed to parse JPEG file: %v", err)
	}
	return context
}

// Set the date for a given photo
func setPhotoDate(filepath string, t time.Time) {

	context := getMediaContext(filepath)

	segmentList := context.(*jpeg.SegmentList)

	rootIb, err := segmentList.ConstructExifBuilder()
	if err != nil {
		fmt.Println("No EXIF; creating it from scratch")

		im, err := exifcommon.NewIfdMappingWithStandard()
		if err != nil {
			log.Fatalf("Failed to create new IFD mapping with standard tags: %v", err)
		}
		ti := exif.NewTagIndex()
		if err := exif.LoadStandardTags(ti); err != nil {
			log.Fatalf("Failed to load standard tags: %v", err)
		}

		rootIb = exif.NewIfdBuilder(im, ti, exifcommon.IfdStandardIfdIdentity, exifcommon.EncodeDefaultByteOrder)
		rootIb.AddStandardWithName("ProcessingSoftware", "")
	}

	// Form our timestamp string
	timestamp := exifcommon.ExifFullTimestampString(t)

	ifdPath := "IFD0"
	setExifTag(rootIb, ifdPath, "DateTime", timestamp)         // Set DateTime
	setExifTag(rootIb, ifdPath, "DateTimeOriginal", timestamp) // Set DateTimeOriginal

	// Update the exif segment.
	if err := segmentList.SetExif(rootIb); err != nil {
		log.Fatalf("Failed to set EXIF to jpeg: %v", err)
	}

	// Write the modified file
	b := new(bytes.Buffer)
	if err := segmentList.Write(b); err != nil {
		log.Fatalf("Failed to create JPEG data: %v", err)
	}

	// Save the file
	if err := ioutil.WriteFile(filepath, b.Bytes(), 0644); err != nil {
		log.Fatalf("Failed to write JPEG file: %v", err)
	}
}
