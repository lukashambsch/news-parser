package utils

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
)

const XMLDir = "xmls"
const ZipDir = "zips"

// Used function from http://stackoverflow.com/questions/20357223/easy-way-to-unzip-file-with-golang
func Unzip(src, dest string) error {
	// create zip reader
	r, err := zip.OpenReader(src)
	if err != nil {
		return err
	}
	defer r.Close()

	os.MkdirAll(dest, 0755)

	// Closure to address file descriptors issue with all the deferred .Close() methods
	extractAndWriteFile := func(f *zip.File) error {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(dest, f.Name)

		// Write directory if file is a directory
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			// Write file if not a directory
			// Create blank file
			os.MkdirAll(filepath.Dir(path), f.Mode())
			// Open file reader
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			// Copy file contents into created file
			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	// iterate over each item in zip
	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func Clean() {
	os.RemoveAll(XMLDir)
	os.RemoveAll(ZipDir)
}
