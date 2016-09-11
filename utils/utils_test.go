package utils_test

import (
	"os"
	"testing"

	"github.com/lukashambsch/news-parser/utils"
)

func TestUnzip(t *testing.T) {
	expectedPath := "./test/test.xml"
	tmpDir := "./test"

	utils.Unzip("test.zip", tmpDir)

	_, err := os.Stat(expectedPath)
	if os.IsNotExist(err) {
		t.Error("Expected ", expectedPath, " to exist.")
	}

	os.RemoveAll(tmpDir)
}
