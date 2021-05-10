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

func getMediaContext(filepath string) riimage.MediaContext {
	parser := jpeg.NewJpegMediaParser()
	context, err := parser.ParseFile(filepath)
	if err != nil {
		log.Panicf("Failed to parse JPEG file: %v", err)
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
			log.Panicf("Failed to create new IFD mapping with standard tags: %v", err)
		}
		ti := exif.NewTagIndex()
		if err := exif.LoadStandardTags(ti); err != nil {
			log.Panicf("Failed to load standard tags: %v", err)
		}

		rootIb = exif.NewIfdBuilder(im, ti, exifcommon.IfdStandardIfdIdentity, exifcommon.EncodeDefaultByteOrder)
		rootIb.AddStandardWithName("ProcessingSoftware", "")
	}

	// Form our timestamp string
	ts := exifcommon.ExifFullTimestampString(t)

	// Set DateTime
	ifdPath := "IFD0"
	if err := setExifTag(rootIb, ifdPath, "DateTime", ts); err != nil {
		log.Panicf("Failed to set tag %v: %v", "DateTime", err)
	}

	// Set DateTimeOriginal
	ifdPath = "IFD/Exif"
	if err := setExifTag(rootIb, ifdPath, "DateTimeOriginal", ts); err != nil {
		log.Panicf("Failed to set tag %v: %v", "DateTimeOriginal", err)
	}

	// Update the exif segment.
	if err := segmentList.SetExif(rootIb); err != nil {
		log.Panicf("Failed to set EXIF to jpeg: %v", err)
	}

	// Write the modified file
	b := new(bytes.Buffer)
	if err := segmentList.Write(b); err != nil {
		log.Panicf("Failed to create JPEG data: %v", err)
	}

	// Save the file
	if err := ioutil.WriteFile(filepath, b.Bytes(), 0644); err != nil {
		log.Panicf("Failed to write JPEG file: %v", err)
	}
}
