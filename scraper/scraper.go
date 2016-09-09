package scraper

import (
	"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

const XMLDir = "xmls"
const ZipDir = "zips"

func Scrape(URL string) {
	var wg sync.WaitGroup
	throttle := make(chan int, 5)
	paths := GetZipLinks(URL)

	for _, path := range paths[0:1] {
		fmt.Printf("Downloading %s\n", path)

		split := strings.Split(path, "/")
		zipPath := split[len(split)-1]

		throttle <- 1
		wg.Add(1)
		go Download(path, fmt.Sprintf("%s/%s", ZipDir, zipPath), &wg, throttle)
	}
	wg.Wait()
}

func Download(path string, zipPath string, wg *sync.WaitGroup, throttle chan int) {
	defer wg.Done()

	os.Mkdir(ZipDir, 0700)
	out, err := os.Create(zipPath)
	if err != nil {
		log.Panic(err)
	}
	defer out.Close()

	response, err := http.Get(path)
	if err != nil {
		log.Panic(err)
	}
	defer response.Body.Close()

	io.Copy(out, response.Body)

	split := strings.Split(zipPath, "/")
	dirName := strings.Replace(split[len(split)-1], ".zip", "", 1)
	outputPath := fmt.Sprintf("%s/%s", XMLDir, dirName)
	Unzip(zipPath, outputPath)

	fmt.Println(fmt.Sprintf("Finished downloading %s\n", path))
	<-throttle
}

func GetZipLinks(URL string) []string {
	var hrefs []string
	doc, _ := goquery.NewDocument(URL)

	doc.Find("table tr td a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")

		if strings.Contains(href, ".zip") {
			hrefs = append(hrefs, fmt.Sprintf("%s%s", URL, href))
		}
	})
	return hrefs
}

func Unzip(src, dest string) error {
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

		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			os.MkdirAll(filepath.Dir(path), f.Mode())
			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
		return nil
	}

	for _, f := range r.File {
		err := extractAndWriteFile(f)
		if err != nil {
			return err
		}
	}

	return nil
}
