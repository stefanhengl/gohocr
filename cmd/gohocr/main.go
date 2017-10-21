/*
gohocr Command Line Interface. Converts tesseract's *.hocr
files to *.json.

Usage: ghocr -f /path/to/hocr/file
*/
package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/stefanhengl/gohocr"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	iPtr := flag.String("f", "", "file")
	flag.Parse()

	if *iPtr == "" {
		log.Println("Use option -f and provide a fully qualified path to a *.hocr file")
		return
	}

	page, err := gohocr.Parse(*iPtr)
	if err != nil {
		log.Println(err.Error())
		return
	}

	pageJSON, _ := json.Marshal(page)
	output := strings.Replace(*iPtr, ".hocr", ".json", 1)
	err = ioutil.WriteFile(output, pageJSON, 0644)
	if err != nil {
		log.Println(err.Error())
		return
	}
	log.Printf("Success. Saved output to %s", output)
}
