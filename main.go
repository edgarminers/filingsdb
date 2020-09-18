package main

import (
	"archive/zip"
	"bufio"
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"eswiac.me/filingsdb/models"
	"github.com/briandowns/spinner"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const BATCH_SIZE = 250

func HandleZip(db *gorm.DB, zipfile string) {
	// Open a zip archive for reading.
	r, err := zip.OpenReader(zipfile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	// Iterate through the files in the archive,
	// printing some of their contents.
	for i, f := range r.File {
		fmt.Printf("Contents of %s:\n", f.Name)
		rc, err := f.Open()
		defer rc.Close()
		if err != nil {
			log.Fatal(err)
		}
		rd := bufio.NewReader(rc)

		// skip header
		rd.ReadString('\n')

		if f.Name == "cal.tsv" {
			var cals = []models.DataCAL{}
			for {
				line, err := rd.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				tokens := strings.Split(line, "\t")
				cal := models.ParseDataCAL(tokens)
				cals = append(cals, cal)
				if len(cals) >= BATCH_SIZE {
					if err := db.Create(&cals).Error; err != nil {
						log.Fatal(err)
					}
					cals = []models.DataCAL{}
				}
			}
			// save the last batch
			if err := db.Create(&cals).Error; err != nil {
				log.Fatal(err)
			}
		}

		if f.Name == "pre.tsv" {
			var pres = []models.DataPRE{}
			for {
				line, err := rd.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				tokens := strings.Split(line, "\t")
				pre := models.ParseDataPRE(tokens)
				pres = append(pres, pre)
				if len(pres) >= BATCH_SIZE {
					if err := db.Create(&pres).Error; err != nil {
						log.Fatal(err)
					}
					pres = []models.DataPRE{}
				}
			}
			// save the last batch
			if err := db.Create(&pres).Error; err != nil {
				log.Fatal(err)
			}
		}

		if f.Name == "ren.tsv" {
			var rens = []models.DataREN{}
			for {
				line, err := rd.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				tokens := strings.Split(line, "\t")
				ren := models.ParseDataREN(tokens)
				rens = append(rens, ren)
				if len(rens) >= BATCH_SIZE {
					if err := db.Create(&rens).Error; err != nil {
						log.Fatal(err)
					}
					rens = []models.DataREN{}
				}
			}
			// save the last batch
			if err := db.Create(&rens).Error; err != nil {
				log.Fatal(err)
			}
		}

		if f.Name == "txt.tsv" {
			var txts = []models.DataTXT{}
			for {
				line, err := rd.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				tokens := strings.Split(line, "\t")
				txt := models.ParseDataTXT(tokens)
				txts = append(txts, txt)
				if len(txts) >= BATCH_SIZE {
					if err := db.Create(&txts).Error; err != nil {
						log.Fatal(err)
					}
					txts = []models.DataTXT{}
				}
			}
			// save the last batch
			if err := db.Create(&txts).Error; err != nil {
				log.Fatal(err)
			}
		}

		if f.Name == "num.tsv" {
			var nums = []models.DataNUM{}
			for {
				line, err := rd.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				tokens := strings.Split(line, "\t")
				num := models.ParseDataNUM(tokens)
				nums = append(nums, num)
				if len(nums) >= BATCH_SIZE {
					if err := db.Create(&nums).Error; err != nil {
						log.Fatal(err)
					}
					nums = []models.DataNUM{}
				}
			}
			// save the last batch
			if err := db.Create(&nums).Error; err != nil {
				log.Fatal(err)
			}
		}

		if f.Name == "dim.tsv" {
			var dims = []models.DataDIM{}
			for {
				line, err := rd.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				tokens := strings.Split(line, "\t")
				dim := models.ParseDataDIM(tokens)
				dims = append(dims, dim)
				if len(dims) >= BATCH_SIZE {
					if err := db.Create(&dims).Error; err != nil {
						log.Fatal(err)
					}
					dims = []models.DataDIM{}
				}
			}
			// save the last batch
			if err := db.Create(&dims).Error; err != nil {
				log.Fatal(err)
			}
		}

		if f.Name == "tag.tsv" {
			var tags = []models.DataTAG{}
			for {
				line, err := rd.ReadString('\n')
				line = strings.TrimRight(line, "\r\n")

				if err == io.EOF {
					break
				}

				if err != nil {
					log.Fatal(err)
				}

				tokens := strings.Split(line, "\t")
				tag := models.ParseDataTAG(tokens)
				tags = append(tags, tag)
				if len(tags) >= BATCH_SIZE {
					if err := db.Create(&tags).Error; err != nil {
						log.Fatal(err)
					}
					tags = []models.DataTAG{}
				}
			}
			// save the last batch
			if err := db.Create(&tags).Error; err != nil {
				log.Fatal(err)
			}
		}

		if f.Name == "sub.tsv" {
			var subs = []models.DataSUB{}
			for {
				// reader.ReadLine does a buffered read up to a line terminator,
				// handles either /n or /r/n, and returns just the line without
				// the /r or /r/n.
				//line, isPrefix, err := bf.ReadLine()
				//...but (http://golang.org/pkg/bufio/#Reader.ReadLine)
				//ReadLine is a low-level line-reading primitive.
				//Most callers should use ReadBytes('\n') or ReadString('\n') instead.
				line, err := rd.ReadString('\n')
				//fmt.Printf("===" + strings.TrimRight(line, "\r\n") + "===")
				line = strings.TrimRight(line, "\r\n")
				// loop termination condition 1:  EOF.
				// this is the normal loop termination condition.
				if err == io.EOF {
					break
				}

				// loop termination condition 2: some other error.
				// Errors happen, so check for them and do something with them.
				if err != nil {
					log.Fatal(err)
				}

				// loop termination condition 3: line too long to fit in buffer
				// without multiple reads.  Bufio's default buffer size is 4K.
				// Chances are if you haven't seen a line terminator after 4k
				// you're either reading the wrong file or the file is corrupt.
				//TODO

				// success.  The variable line is now a byte slice based on on
				// bufio's underlying buffer.  This is the minimal churn necessary
				// to let you look at it, but note! the data may be overwritten or
				// otherwise invalidated on the next read.  Look at it and decide
				// if you want to keep it.  If so, copy it or copy the portions
				// you want before iterating in this loop.  Also note, it is a byte
				// slice.  Often you will want to work on the data as a string,
				// and the string type conversion (shown here) allocates a copy of
				// the data.  It would be safe to send, store, reference, or otherwise
				// hold on to this string, then continue iterating in this loop.
				tokens := strings.Split(line, "\t")
				subs = append(subs, models.ParseDataSUB(tokens))
				//fmt.Println("> " + tokens[len(tokens)-1])
				if len(subs) >= BATCH_SIZE {
					if err := db.Create(&subs).Error; err != nil {
						log.Fatalf("line [%s] - %s", line, err)
					}
					subs = []models.DataSUB{}
				}
			}
			if err := db.Create(&subs).Error; err != nil {
				log.Fatalf("%d - %s", i, err.Error())
			}
		}

	}
}

func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Minute,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      true,          // Disable color
		},
	)
	db, err := gorm.Open(sqlite.Open("filings.db"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
	}

	db.Exec("DROP TABLE IF EXISTS data_subs;")
	db.Exec("DROP TABLE IF EXISTS data_tags;")
	db.Exec("DROP TABLE IF EXISTS data_dims;")
	db.Exec("DROP TABLE IF EXISTS data_nums;")
	db.Exec("DROP TABLE IF EXISTS data_txts;")
	db.Exec("DROP TABLE IF EXISTS data_pres;")
	db.Exec("DROP TABLE IF EXISTS data_rens;")
	db.Exec("DROP TABLE IF EXISTS data_cals;")

	db.AutoMigrate(
		&models.DataSUB{},
		&models.DataTAG{},
		&models.DataDIM{},
		&models.DataNUM{},
		&models.DataTXT{},
		&models.DataPRE{},
		&models.DataREN{},
		&models.DataCAL{},
	)

	if len(os.Args) == 2 {
		ProcessURL(db, os.Args[1])
		os.Exit(0)
	}

	links, err := GetArchivesURLs()
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("Found %v", links)
	for _, url := range links {
		ProcessURL(db, url)
	}
}

func ProcessURL(db *gorm.DB, url string) {
	err := os.Mkdir(".cache", 0755)
	h := md5.New()
	io.WriteString(h, url)
	md5NameHash := fmt.Sprintf("%x", h.Sum(nil))
	filePath := ".cache/" + md5NameHash

	_, err = os.Stat(filePath)
	var file *os.File
	defer file.Close()

	if os.IsNotExist(err) {
		file, err = os.Create(filePath)
		if err != nil {
			panic(err)
		}
		log.Printf("Downloading %v ...", url)
		err = DownloadFile(url, file.Name())
		if err != nil {
			log.Fatal(err)
		}
	} else {
		file, err = os.Open(filePath)
		if err != nil {
			panic(err)
		}
	}

	log.Print("Processing, please be patient")
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond) // Build our new spinner
	s.Start()                                                    // Start the spinner
	HandleZip(db, file.Name())
	s.Stop()
	log.Println()
}
