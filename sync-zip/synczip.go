package synczip

import (
	"archive/zip"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"
)

// Zipper make .zip file from urls
func Zipper(urls []string) {
	var wg sync.WaitGroup
	wg.Add(len(urls))
	for _, url := range urls {
		go func(url string) {
			defer wg.Done()
			if _, err := download(url); err != nil {
				log.Fatal(err)
			}
		}(url)
	}
	wg.Wait()

	filenames, err := filepath.Glob("*.png")
	if err != nil {
		log.Fatal(err)
	}

	err = writeZip("icons.zip", filenames)
	if err != nil {
		log.Fatal(err)
	}
}

func writeZip(outFilename string, filenames []string) error {
	outf, err := os.Create(outFilename)
	if err != nil {
		return err
	}
	zw := zip.NewWriter(outf)
	for _, filename := range filenames {
		w, err := zw.Create(filename)
		if err != nil {
			return err
		}
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
		_, err = io.Copy(w, f)
		if err != nil {
			return err
		}
	}
	return zw.Close()
}

func download(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	filename, err := urlToFilename(url)
	if err != nil {
		return "", err
	}
	f, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer f.Close()
	_, err = io.Copy(f, resp.Body)
	return filename, err
}

func urlToFilename(rawurl string) (string, error) {
	url, err := url.Parse(rawurl)
	if err != nil {
		return "", err
	}
	return filepath.Base(url.Path), nil
}
