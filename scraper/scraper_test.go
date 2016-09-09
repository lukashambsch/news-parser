package scraper_test

import (
	"os"
	"testing"

	"github.com/lukashambsch/news-parser/scraper"
)

func TestUnzip(t *testing.T) {
	expectedPath := "./test/test.xml"
	tmpDir := "./test"

	scraper.Unzip("test.zip", tmpDir)

	_, err := os.Stat(expectedPath)
	if os.IsNotExist(err) {
		t.Error("Expected ", expectedPath, " to exist.")
	}

	os.RemoveAll(tmpDir)
}
