package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"eswiac.me/filingsdb/models"
	"github.com/antchfx/htmlquery"
	"github.com/briandowns/spinner"
	"github.com/dustin/go-humanize"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetArchivesURLs() []string {
	secGov := "https://www.sec.gov"
	url := secGov + "/dera/data/financial-statement-and-notes-data-set.html"
	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Fatal(err)
	}
	list := htmlquery.Find(doc, "//a/@href")
	links := []string{}
	for _, n := range list {
		link := htmlquery.SelectAttr(n, "href") // output @href value
		if strings.HasSuffix(link, "_notes.zip") {
			links = append(links, secGov+link)
		}
	}
	return links
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

func (d *Downloader) DownloadFile(url string, filepath string) {

	// Create the file with .tmp extension, so that we won't overwrite a
	// file until it's downloaded fully
	out, err := os.Create(filepath + ".tmp")
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Create our bytes counter and pass it to be used alongside our writer
	counter := &WriteCounter{}
	_, err = io.Copy(out, io.TeeReader(resp.Body, counter))
	if err != nil {
		log.Fatal(err)
	}

	// The progress use the same line so print a new line once it's finished downloading
	fmt.Println()

	// Rename the tmp file back to the original file
	err = os.Rename(filepath+".tmp", filepath)
	if err != nil {
		log.Fatal(err)
	}
}

type Downloader struct {
	db       *gorm.DB
	yearUrls []string
	year     string
}

func dbName(year string) string {
	return fmt.Sprintf("filings_%v.db", year)
}

func New(year string) *Downloader {
	yearInt, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if yearInt < 2009 || yearInt > 2100 {
		log.Fatalf("Filings are not available before 2009 and after 2100")
	}
	yearUrls := []string{} // 4 qtr in a year
	for _, url := range GetArchivesURLs() {
		if strings.Contains(url, year) {
			yearUrls = append(yearUrls, url)
		}
	}
	if len(yearUrls) == 0 {
		log.Fatalf("Couldn't find any filings from sec.gov filed in %v", year)
	}

	_, err = os.Stat(dbName(year))
	if !os.IsNotExist(err) {
		log.Fatalf("filings database already exists. This script does not perform differential updates. Please rename or move %v in order to rebuild the filings database.", dbName(year))
	}
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Minute,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open(dbName(year)+"?_journal_mode=WAL"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(
		&models.DataSUB{},
		&models.DataTAG{},
		&models.DataDIM{},
		&models.DataNUM{},
		&models.DataTXT{},
		&models.DataPRE{},
		&models.DataREN{},
		&models.DataCAL{},
		&models.DataTicker{},
	)
	return &Downloader{yearUrls: yearUrls, db: db, year: year}
}
func (d Downloader) Start() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Start()
	fmt.Println("Processing started, please be patient. This may take a while!")
	d.downloadTickers()

	for _, url := range d.yearUrls {
		d.handle(url)
	}

	s.Stop()
	fmt.Printf("Processing complete. You can now open your %v filings database with `sqlite3 %v`", d.year, dbName((d.year)))
	fmt.Println()
}

func (d Downloader) handle(url string) {
	tmpFile, err := ioutil.TempFile("", "")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	d.DownloadFile(url, tmpFile.Name())
	ExtractFromZip(d.db, tmpFile.Name())
}

func (d Downloader) downloadTickers() {
	tickersList := models.DataTickers{}
	// Get the data
	resp, err := http.Get("https://www.sec.gov/files/company_tickers.json")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &tickersList)
	if err != nil {
		log.Fatal(err)
	}
	tickers := []models.DataTicker{}
	for _, ticker := range tickersList {
		ticker.CikString = strconv.Itoa(ticker.Cik)
		tickers = append(tickers, ticker)
		if len(tickers) > BATCH_SIZE {
			if err := d.db.Create(&tickers).Error; err != nil {
				log.Fatal(err)
			}
			tickers = []models.DataTicker{}
		}
	}
	// last batch
	if err := d.db.Create(&tickers).Error; err != nil {
		log.Fatal(err)
	}
}
