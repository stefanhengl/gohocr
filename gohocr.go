/*
Package gohocr parses hocr files.
*/
package gohocr

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"os"
)

// Page represents one hocr file
type Page struct {
	Words []Word `xml:"body>div>div>p>span>span"`
}

// Word represents a single word within an hocr file
// <span class='ocrx_word' id='word_1_11' title='bbox 572 568 641 684; x_wconf 57' lang='eng' dir='ltr'>Foo</span>
type Word struct {
	Lang      string `xml:"lang,attr"`
	Direction string `xml:"dir,attr"`
	Title     string `xml:"title,attr"`
	ID        string `xml:"id,attr"`
	Class     string `xml:"class,attr"`
	Content   string `xml:",innerxml"`
}

// Parse takes either a string, *os.File, or []byte and returns a Page object
func Parse(value interface{}) (Page, error) {
	var byteValue []byte
	var err error
	switch str := value.(type) {
	case string:
		xmlFile, err := os.Open(str)
		if err != nil {
			return Page{}, err
		}
		defer xmlFile.Close()
		byteValue, err = ioutil.ReadAll(xmlFile)
		if err != nil {
			return Page{}, err
		}
	case *os.File:
		byteValue, err = ioutil.ReadAll(str)
		if err != nil {
			return Page{}, err
		}
	case []byte:
		byteValue = str
	default:
		return Page{}, errors.New("Invalid input for Parse. Submit either a string, *os.File, or []byte")
	}

	var page Page
	err = xml.Unmarshal(byteValue, &page)
	if err != nil {
		return Page{}, err
	}
	return page, nil
}
