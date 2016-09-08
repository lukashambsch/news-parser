package parser

import (
	"encoding/xml"
	//"fmt"
	"os"

	"github.com/lukashambsch/news-parser/store"
)

func ParseXML(path string) store.Document {
	var doc store.Document
	f, _ := os.Open(path)
	decoder := xml.NewDecoder(f)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}

		switch se := t.(type) {
		case xml.StartElement:
			if se.Name.Local == "document" {
				decoder.DecodeElement(&doc, &se)
			}
		}
	}

	//fmt.Printf("%#v", doc)
	return doc
}
