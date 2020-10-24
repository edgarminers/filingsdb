package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if len(os.Args) != 2 {
		fmt.Println("Usage: filingsdb <year>")
		os.Exit(-1)
	}
	downloader := New(os.Args[1])
	downloader.Start()
}
