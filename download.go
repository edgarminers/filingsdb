package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/dustin/go-humanize"
)

const SEC_GOV_HOST = "https://www.sec.gov"

func GetArchivesURLs() ([]string, error) {
	url := SEC_GOV_HOST + "/dera/data/financial-statement-and-notes-data-set.html"
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		return nil, err
	}
	list := htmlquery.Find(doc, "//a/@href")
	links := []string{}
	for _, n := range list {
		link := htmlquery.SelectAttr(n, "href") // output @href value
		if strings.HasSuffix(link, "_notes.zip") {
			links = append(links, SEC_GOV_HOST+link)
		}
	}
	return links, nil
}

func (wc *WriteCounter) Write(p []byte) (int, error) {
	n := len(p)
	wc.Total += uint64(n)
	wc.PrintProgress()
	return n, nil
}

type WriteCounter struct {
	Total uint64
}

// PrintProgress prints the progress of a file write
func (wc WriteCounter) PrintProgress() {
	// Clear the line by using a character return to go back to the start and remove
	// the remaining characters by filling it with spaces
	fmt.Printf("\r%s", strings.Repeat(" ", 50))

	// Return again and print current status of download
	// We use the humanize package to print the bytes in a meaningful way (e.g. 10 MB)
	fmt.Printf("\rDownloading... %s complete", humanize.Bytes(wc.Total))
}

func DownloadFile(url string, filepath string) error {

	// Create the file with .tmp extension, so that we won't overwrite a
	// file until it's downloaded fully
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create our bytes counter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		return err
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Println()

	// Rename the tmp file back to the original file
	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		return err
	}
	return nil
}
