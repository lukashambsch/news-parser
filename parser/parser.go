package parser

import (
	"encoding/xml"
	"os"

	"github.com/lukashambsch/news-parser/models"
)

func ParseXML(path string) models.Post {
	var post models.Post
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
				decoder.DecodeElement(&post, &se)
			}
		}
	}

	return post
}
