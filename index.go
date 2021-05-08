package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/Jeffail/gabs"
)

func main() {
	fmt.Println("Starting...")
	getFiles()
	readFile("test")
}

func readFile(filename string) {
	jsonParsed, err := gabs.ParseJSON([]byte(`{
		"outter":{
			"inner":{
				"value1":10,
				"value2":22
			},
			"alsoInner":{
				"value1":20,
				"array1":[
					30, 40
				]
			}
		}
	}`))
	if err != nil {
		panic(err)
	}

	fmt.Print(jsonParsed)
}

func getFiles() {
	err := filepath.Walk("./photos_and_videos/album",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			fmt.Println(path, info.Size())
			return nil
		})
	if err != nil {
		log.Println(err)
	}
}
